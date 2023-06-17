package handler

import (
	"context"
	"encoding/json"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gdamore/tcell/v2"
	"github.com/go-redis/redis/v8"
	"github.com/hashicorp/consul/api"
	"github.com/streadway/amqp"
	"google.golang.org/protobuf/proto"
	"log"
	"net"
	"strconv"
	"sync"
	"tank_war/server/cmd/game/config"
	pb "tank_war/server/cmd/game/handler/pb/quic"
	"tank_war/server/shared/consts"
	"tank_war/server/shared/kitex_gen/base"
	"time"
)

var roomMutex sync.Mutex

type Room struct {
	id         int64
	name       string
	clients    map[int64]*Client
	unregistry chan *Client
	registry   chan *Client
	broadcast  chan *pb.Action
	handler    *Handler
	maxPlayer  int
	gameOver   chan struct{}
}

var rooms = make(map[int64]*Room)

func getRoom(req *pb.JoinRoomReq) *Room {
	if v, ok := rooms[req.RoomId]; ok {
		return v
	}
	r := &Room{
		id:         req.RoomId,
		clients:    make(map[int64]*Client),
		unregistry: make(chan *Client),
		registry:   make(chan *Client),
		broadcast:  make(chan *pb.Action),
		handler:    NewHandler(),
		maxPlayer:  int(req.MaxPlayer),
		name:       req.RoomName,
	}
	rooms[req.RoomId] = r
	go r.Run()
	klog.Infof("create room:", req.RoomId)
	if len(rooms) > 0 {
		updateServiceMeta("busy")
	}
	return r
}

func (r *Room) Run() {
	go func() {
		for {
			actions := r.handler.UpdateStatus()
			for _, action := range actions {
				r.broadcast <- action
			}
			time.Sleep(100 * time.Millisecond) //TODO : 修改游戏帧率
		}
	}()
	for {
		select {
		case c := <-r.registry:
			r.Registry(c)
		case c := <-r.unregistry:
			r.Unregistry(c)
		case action := <-r.broadcast:
			r.Broadcast(action)
		}
	}
}

func (r *Room) Route(data []byte) {
	act := &pb.Action{}
	err := proto.Unmarshal(data, act)
	if err != nil {
		log.Println(err)
	}

	switch act := act.Type.(type) {
	case *pb.Action_TankMove:
		r.handler.TankMove(act)
	case *pb.Action_NewBullet:
		r.handler.NewBullet(act)
	default:
	}
}

func (r *Room) Broadcast(action *pb.Action) {

	//TODO: broadcast data to all clients concurrently
	for _, c := range r.clients {
		c.send <- action
	}
}

func (r *Room) Unregistry(c *Client) {
	if _, ok := r.clients[c.id]; ok {
		delete(r.clients, c.id)
		close(c.send)
	}

	if len(r.clients) == 0 {
		if _, ok := rooms[c.room.id]; !ok {
			return
		}
		delete(rooms, c.room.id)
		klog.Infof("delete room:", c.room.id)
		if len(rooms) < 1 {
			updateServiceMeta("free")
		}
		err := config.Rdb.HDel(context.Background(), consts.Room, r.name).Err()
		if err != nil {
			klog.Infof("delete room from redis error :%v", err)
		}
		r.addToMq()
	}
}

func (r *Room) Registry(c *Client) {

	r.clients[c.id] = c
	//frame
	color := uint64(tcell.ColorGreen + tcell.Color(len(r.clients)))
	r.handler.NewTank(c.name, c.id, color)
	c.send <- r.handler.GetRockList()
	if len(r.clients) == r.maxPlayer {
		r.handler.status = consts.GameStart
		err := config.Rdb.HDel(context.Background(), consts.Room, r.name).Err()
		if err != nil {
			klog.Infof("delete room from redis error :%v", err)
		}
	}

	roomInfo := &base.Room{}
	//get room from redis
	jsonData, err := config.Rdb.HGet(context.Background(), consts.Room, r.name).Bytes()
	if err != nil {
		if err == redis.Nil {
			return
		}
		klog.Infof("get room from redis error :%v", err)
	}
	err = json.Unmarshal(jsonData, roomInfo)
	if err != nil {
		klog.Infof("unmarshal room error :%v", err)
	}
	roomInfo.CurrentPlayer++
	jsonData, err = json.Marshal(roomInfo)
	if err != nil {
		klog.Infof("marshal room error :%v", err)
	}
	config.Rdb.HSet(context.Background(), consts.Room, r.name, jsonData)
}

func (r *Room) GetIdList() []int64 {
	var l []int64
	for _, v := range r.clients {
		l = append(l, v.id)
	}
	log.Println("GetIdList:", l)
	return l
}

func updateServiceMeta(status string) {
	cfg := api.DefaultConfig()
	cfg.Address = net.JoinHostPort(
		config.GlobalConsulConfig.Host,
		strconv.Itoa(config.GlobalConsulConfig.Port))
	client, err := api.NewClient(cfg)
	if err != nil {
		klog.Infof("updatetServiceMeta new consul client error :%v", err)
		return
	}

	check := &api.AgentServiceCheck{
		TTL:                            consts.ConsulCheckTTL,
		DeregisterCriticalServiceAfter: consts.ConsulCheckDeregisterCriticalServiceAfter,
		Timeout:                        consts.ConsulCheckTimeout,
	}

	r := &api.AgentServiceRegistration{
		ID:      "game_srv:" + net.JoinHostPort(config.GlobalServerConfig.Host, strconv.Itoa(config.GlobalServerConfig.Port)),
		Address: config.GlobalServerConfig.Host,
		Name:    "game_srv",
		Port:    config.GlobalServerConfig.Port,
		Check:   check,
		Meta:    map[string]string{"status": status},
	}
	err = client.Agent().ServiceRegister(r)

	if err != nil {
		klog.Infof("updateServiceMeta error :%v", err)
		return
	}

	klog.Infof("updateServiceMeta success")

}

//
//func GetRoomList() []*base.Room {
//	var roomList []*base.Room
//	roomMap, err := config.Rdb.HGetAll(context.Background(), consts.Room).Result()
//	if err != nil {
//		klog.Infof("get room list error :%v", err)
//		return roomList
//	}
//	for _, v := range roomMap {
//		room := &base.Room{}
//		err := json.Unmarshal([]byte(v), room)
//		if err != nil {
//			klog.Infof("unmarshal room error :%v", err)
//			continue
//		}
//		roomList = append(roomList, room)
//	}
//	return roomList
//}

func (r *Room) addToMq() {
	for _, v := range r.handler.game.TankBucket {
		var u = &User{
			ID:       v.Id,
			Username: v.Name,
			Kill:     v.Kill,
		}
		if v.IsDead {
			u.Death = 1
		}
		data, err := json.Marshal(u)
		if err != nil {
			klog.Infof("marshal user error :%v", err)
			continue
		}
		err = config.MqChan.Publish(
			"",
			"user_data",
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        data,
			})
		if err != nil {
			klog.Infof("publish message error :%v", err)
		}
		klog.Infof("publish message success")
	}
}

type User struct {
	ID       int64  `gorm:"column:id;primary_key" json:"id"`
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Kill     int32  `gorm:"column:kill" json:"kill"`
	Death    int32  `gorm:"column:death" json:"death"`
}
