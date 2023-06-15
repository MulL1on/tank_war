package room

import (
	"bytes"
	"encoding/binary"
	"github.com/quic-go/quic-go"
	"google.golang.org/protobuf/proto"
	"log"
	pb "tank_war/server/cmd/game/handler/pb/quic"
	"time"
)

type Client struct {
	room   *Room
	id     int32
	stream quic.Stream
	send   chan *pb.Action
}

func NewClient(stream quic.Stream, id int32) {

	room := getRoom(1)
	c := &Client{
		id:     id,
		room:   room,
		stream: stream,
		send:   make(chan *pb.Action),
	}

	room.registry <- c
	go c.WritePump()
	go c.ReadPump()
}

func (c *Client) ReadPump() {
	defer func() {
		c.room.unregistry <- c
		c.stream.Close()
	}()

	for {
		data := make([]byte, 1024)
		_, err := c.stream.Read(data)
		if err != nil {
			c.room.unregistry <- c
			c.stream.Close()
			break
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

		// 路由到对应的处理函数
		c.room.Route(data)
	}
}

func (c *Client) WritePump() {
	defer func() {
		c.stream.Close()
	}()

	for {
		select {
		case action, ok := <-c.send:
			if !ok {
				c.stream.Write([]byte("close"))
				return
			}
			data, err := proto.Marshal(action)
			if err != nil {
				log.Println(err)
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

			n, err := c.stream.Write(buffer.Bytes())
			if err != nil {
				log.Println(err)
			}
			if n != len(buffer.Bytes()) {
				log.Println("send length error")
			}
			time.Sleep(time.Millisecond * 1)
			//TODO: 为什么加了这个log就不会出现数据丢失的情况
			//log.Println("send length:", len(data), "total length:", len(buffer.Bytes()))
		}
	}
}
