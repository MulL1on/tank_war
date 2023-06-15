package handler

import (
	"bytes"
	"encoding/binary"
	"github.com/quic-go/quic-go"
	"google.golang.org/protobuf/proto"
	"log"
	"tank_war/client/game"
	pb "tank_war/client/handler/pb/quic"
)

type Client struct {
	id      int64
	stream  quic.Stream
	exit    chan struct{}
	handler *Handler
	send    chan *pb.Action
}

func NewClient(stream quic.Stream, id int64) {
	// 初始化游戏
	game.NewGame(id)

	var c = &Client{
		id:     id,
		stream: stream,
		exit:   make(chan struct{}),
		send:   make(chan *pb.Action),
	}
	c.handler = NewHandler(c)
	go c.handler.Listen()
	go c.WriteMessage()
	go c.ReadMessage()
}

func (c *Client) ReadMessage() {
	defer c.stream.Close()
	for {
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

func (c *Client) WriteMessage() {
	defer c.stream.Close()
	for {
		select {
		case <-c.exit:
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
	switch act := act.Type.(type) {
	case *pb.Action_GetBulletList:
		c.handler.GetBulletList(act)
	case *pb.Action_GetRockList:
		c.handler.GetRockList(act)
	case *pb.Action_GetTankList:
		c.handler.GetTankList(act)
	case *pb.Action_GetExplosionList:
		c.handler.GetExplosionList(act)
	}
}
