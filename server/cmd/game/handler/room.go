package handler

import (
	"context"
	"encoding/json"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hashicorp/consul/api"
	"google.golang.org/protobuf/proto"
	"log"
	"net"
	"strconv"
	"tank_war/server/cmd/game/config"
	pb "tank_war/server/cmd/game/handler/pb/quic"
	"tank_war/server/shared/consts"
	"tank_war/server/shared/kitex_gen/base"
	"time"
)

type Room struct {
	name       string
	clients    map[int64]*Client
	unregistry chan *Client
	registry   chan *Client
	broadcast  chan *pb.Action
	handler    *Handler
	maxPlayer  int
}

var rooms = make(map[int64]*Room)

func getRoom(req *pb.JoinRoomReq) *Room {
	if v, ok := rooms[req.RoomId]; ok {
		return v
	}
	r := &Room{
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
	if len(rooms) > 4 {
		updateServiceMeta()
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
		case data := <-r.broadcast:
			r.Broadcast(data)
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
}

func (r *Room) Registry(c *Client) {
	r.clients[c.id] = c
	//frame
	r.handler.NewTank(c.id)
	c.send <- r.handler.GetRockList()
	if len(r.clients) == r.maxPlayer {
		r.handler.status = consts.GameStart
	}

	roomInfo := &base.Room{}
	//get room from redis
	jsonData, err := config.Rdb.HGet(context.Background(), consts.Room, r.name).Bytes()
	if err != nil {
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

func updateServiceMeta() {
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
		ID:      "game_srv:" + config.GlobalConsulConfig.Host + strconv.Itoa(config.GlobalServerConfig.Port),
		Address: config.GlobalServerConfig.Host,
		Name:    "game_srv",
		Port:    config.GlobalServerConfig.Port,
		Check:   check,
		Meta:    map[string]string{"status": "busy"},
	}
	err = client.Agent().ServiceRegister(r)

	if err != nil {
		klog.Infof("updateServiceMeta error :%v", err)
		return
	}

	klog.Infof("updateServiceMeta success")

}
