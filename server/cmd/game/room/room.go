package room

import (
	"google.golang.org/protobuf/proto"
	"log"
	"tank_war/game/handler"
	pb "tank_war/quic/handler/pb/quic"
	"time"
)

type Room struct {
	clients    map[int32]*Client
	unregistry chan *Client
	registry   chan *Client
	broadcast  chan *pb.Action
	handler    *handler.Handler
}

var rooms = make(map[int32]*Room)

func getRoom(roomId int32) *Room {
	if v, ok := rooms[roomId]; ok {
		return v
	}
	r := &Room{
		clients:    make(map[int32]*Client),
		unregistry: make(chan *Client),
		registry:   make(chan *Client),
		broadcast:  make(chan *pb.Action),
		handler:    handler.NewMessageManager(),
	}
	rooms[roomId] = r
	go r.Run()
	log.Println("create room:", roomId)
	return r
}

func (r *Room) Run() {
	go func() {
		for {
			actions := r.handler.UpdateStatus()
			for _, action := range actions {
				r.broadcast <- action
			}
			time.Sleep(100 * time.Millisecond)
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
}

func (r *Room) GetIdList() []int32 {
	var l []int32
	for _, v := range r.clients {
		l = append(l, v.id)
	}
	log.Println("GetIdList:", l)
	return l
}
