package handler

import (
	"bytes"
	"encoding/binary"
	"github.com/gdamore/tcell/v2"
	"github.com/quic-go/quic-go"
	"google.golang.org/protobuf/proto"
	"log"
	"tank_war/client/game"
	pb "tank_war/client/handler/pb/quic"
)

type Client struct {
	id      int64
	stream  quic.Stream
	Exit    chan struct{}
	handler *Handler
	send    chan *pb.Action
}

func NewClient(stream quic.Stream, id int64, screen tcell.Screen) *Client {
	// 初始化游戏
	game.NewGame(id)

	var c = &Client{
		id:     id,
		stream: stream,
		Exit:   make(chan struct{}),
		send:   make(chan *pb.Action),
	}
	c.handler = NewHandler(c, screen)
	go c.WriteMessage()
	go c.ReadMessage()
	go c.handler.Listen()
	return c
}

func (c *Client) ReadMessage() {
	defer func() {
		c.stream.Close()
		c.Exit <- struct{}{}
	}()
	for {
		select {
		case <-c.Exit:
			return
		default:
			data := make([]byte, 1024)
			_, err := c.stream.Read(data)
			if err != nil {
				log.Println(err)
				return
			}
			buffer := bytes.NewBuffer(data)
			// read data length
			var length int32
			if err = binary.Read(buffer, binary.BigEndian, &length); err != nil {
				log.Println(err)
			}
			// read data content
			data = make([]byte, length)
			if err = binary.Read(buffer, binary.BigEndian, &data); err != nil {
				log.Println(err)
			}

			c.Route(data)
		}
	}
}

func (c *Client) WriteMessage() {
	defer func() {
		c.stream.Close()
		c.Exit <- struct{}{}
	}()
	for {
		select {
		case <-c.Exit:
			return
		case action := <-c.send:
			data, err := proto.Marshal(action)
			if err != nil {
				log.Println(err)
				return
			}
			buffer := bytes.NewBuffer([]byte{})

			// write data length
			if err = binary.Write(buffer, binary.BigEndian, int32(len(data))); err != nil {
				log.Println(err)
			}
			// write data content
			if err = binary.Write(buffer, binary.BigEndian, data); err != nil {
				log.Println(err)
			}

			_, err = c.stream.Write(buffer.Bytes())
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func (c *Client) Route(data []byte) {
	act := &pb.Action{}
	err := proto.Unmarshal(data, act)
	if err != nil {
		log.Println(err)
	}

	// 加锁保护

	switch act := act.Type.(type) {
	case *pb.Action_GetBulletList:
		c.handler.GetBulletList(act)
	case *pb.Action_GetRockList:
		c.handler.GetRockList(act)
	case *pb.Action_GetTankList:
		c.handler.GetTankList(act)
	case *pb.Action_GetExplosionList:
		c.handler.GetExplosionList(act)
	case *pb.Action_GameOver:
		c.Exit <- struct{}{}
	}
}
