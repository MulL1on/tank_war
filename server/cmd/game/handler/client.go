package handler

import (
	"bytes"
	"encoding/binary"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/quic-go/quic-go"
	"google.golang.org/protobuf/proto"
	pb "tank_war/server/cmd/game/handler/pb/quic"
	"tank_war/server/shared/consts"
	"time"
)

type Client struct {
	name   string
	room   *Room
	id     int64
	stream quic.Stream
	send   chan *pb.Action
}

func NewClient(stream quic.Stream, req *pb.JoinRoomReq) {

	room := getRoom(req)
	if _, ok := room.clients[req.PlayerId]; ok {
		stream.Close()
		return
	}
	c := &Client{
		id:     req.PlayerId,
		room:   room,
		stream: stream,
		send:   make(chan *pb.Action),
		name:   req.PlayerName,
	}

	room.registry <- c
	go c.writePump()
	go c.readPump()
}

func (c *Client) readPump() {
	defer func() {
		c.room.unregister <- c
		c.stream.Close()
	}()

	for {
		data := make([]byte, 1024)
		_, err := c.stream.Read(data)
		if err != nil {
			c.room.unregister <- c
			c.stream.Close()
			break
		}
		if c.room.handler.status != consts.GameStart {
			continue
		}

		buffer := bytes.NewBuffer(data)
		// read data length
		var length int32
		if err = binary.Read(buffer, binary.BigEndian, &length); err != nil {
			klog.Infof("read data length error: %v", err)
		}

		// read data content
		data = make([]byte, length)
		if err = binary.Read(buffer, binary.BigEndian, &data); err != nil {
			klog.Infof("read data content error: %v", err)
		}

		// 路由到对应的处理函数
		c.room.Route(data)
	}
}

func (c *Client) writePump() {
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
				klog.Infof("proto marshal error: %v", err)
			}

			buffer := bytes.NewBuffer([]byte{})

			// write data length
			if err = binary.Write(buffer, binary.BigEndian, int32(len(data))); err != nil {
				klog.Infof("write data length error")
			}
			// write data content
			if err = binary.Write(buffer, binary.BigEndian, data); err != nil {
				klog.Infof("write data content error: %v", err)
			}

			n, err := c.stream.Write(buffer.Bytes())
			if err != nil {
				klog.Infof("write data error: %v", err)
			}
			if n != len(buffer.Bytes()) {
				klog.Infof("send length error: %v", err)
			}
			time.Sleep(time.Millisecond * 1)
		}
	}
}
