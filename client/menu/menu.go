package menu

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/eiannone/keyboard"
	"github.com/gdamore/tcell/v2"
	"github.com/quic-go/quic-go"
	"google.golang.org/protobuf/proto"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"tank_war/client/consts"
	"tank_war/client/game"
	"tank_war/client/handler"
	pb "tank_war/client/handler/pb/quic"
	"tank_war/server/cmd/api/biz/model/room"
	"tank_war/server/shared/kitex_gen/user"
	"time"
)

var token string

func register() {
	fmt.Println()
	fmt.Println("please enter username:")
	var username string
	fmt.Scanln(&username)
	fmt.Println("please enter password:")
	var password string
	fmt.Scanln(&password)

	URL := getUrl("/user/register") // 设置创建房间的URL
	// 构造请求的参数
	var args protocol.Args
	args.Set("username", username)
	args.Set("password", password)

	c, err := client.NewClient()
	if err != nil {
		fmt.Println("new client failed:", err)
	}
	status, body, err := c.Post(context.Background(), nil, URL, &args)
	if err != nil {
		fmt.Println("do request failed:", err)
	}

	if status != http.StatusOK {
		fmt.Println("register failed: " + string(body))
	} else {
		fmt.Println(string(body))
	}
	pressAnyButtonToContinue()

}

func login() {
	fmt.Println()
	if token != "" {
		fmt.Println("you have already login")
		pressAnyButtonToContinue()
		return
	}

	var username string
	fmt.Println("please enter username:")
	fmt.Scanln(&username)
	fmt.Println("please enter password:")
	var password string
	fmt.Scanln(&password)

	URL := getUrl("/user/login") // 设置创建房间的URL

	c, err := client.NewClient()
	if err != nil {
		fmt.Println("new client failed:", err)
	}

	// 构造请求的参数
	ags := protocol.Args{}
	ags.Set("username", username)
	ags.Set("password", password)

	// 发送POST请求
	status, body, err := c.Post(context.Background(), nil, URL, &ags)
	if err != nil {
		fmt.Println("do request failed:", err)
	}

	if status != http.StatusOK {
		fmt.Println("login failed " + string(body))
	} else {
		var res = &user.LoginResp{}
		//unmarshal
		err = json.Unmarshal(body, res)
		if err != nil {
			fmt.Println("unmarshal failed:", err)
		}
		token = res.Token
		fmt.Println("login success")
	}
	pressAnyButtonToContinue()
}

func createRoom() {
	fmt.Println()

	if token == "" {
		fmt.Println("please login first")
		pressAnyButtonToContinue()
		return
	}
	var roomName string
	fmt.Println("please enter room name:")
	fmt.Scanln(&roomName)
	fmt.Println("please enter room capacity:")
	var capacity int
	fmt.Scanln(&capacity)

	URL := getUrl("/room") // 设置创建房间的URL
	// 构造请求的参数
	req, resp := protocol.AcquireRequest(), protocol.AcquireResponse()

	c, err := client.NewClient()
	if err != nil {
		fmt.Println("new client failed:", err)
	}
	req.SetHeader("Authorization", token)

	req.SetRequestURI(URL)
	req.SetMethod("POST")
	req.SetFormData(map[string]string{
		"name":       roomName,
		"max_player": strconv.Itoa(capacity),
	})

	// 发送POST请求
	err = c.DoTimeout(context.Background(), req, resp, time.Second*5)

	if err != nil {
		fmt.Println("do request failed:", err)
	}

	if resp.StatusCode() != http.StatusOK {
		fmt.Println("create room failed: " + string(resp.Body()))
	} else {
		var res = &room.JoinRoomResp{}
		//unmarshal
		err = json.Unmarshal(resp.Body(), res)
		if err != nil {
			fmt.Println("unmarshal failed:", err)
		}

		startTheGame(res)
	}
	pressAnyButtonToContinue()
}

func joinRoom() {
	fmt.Println()
	if token == "" {
		fmt.Println("please login first")
		pressAnyButtonToContinue()
		return
	}
	var roomName string
	fmt.Println("please enter room name:")
	fmt.Scanln(&roomName)

	URL := getUrl("/room/join") // 设置创建房间的URL

	c, err := client.NewClient()
	if err != nil {
		fmt.Println("new client failed:", err)
	}
	req, resp := protocol.AcquireRequest(), protocol.AcquireResponse()
	req.SetRequestURI(URL)
	req.Header.Set("Authorization", token)
	req.SetMethod("POST")
	req.SetFormData(map[string]string{
		"name": roomName,
	})
	err = c.DoTimeout(context.Background(), req, resp, time.Second*5)
	if err != nil {
		fmt.Println("do request failed:", err)
	}
	status := resp.StatusCode()
	body := resp.Body()

	//TODO: get room info from agent server && start quic
	if status != http.StatusOK {
		fmt.Println("join room failed " + string(body))
	} else {
		var res = &room.JoinRoomResp{}
		//unmarshal
		err = json.Unmarshal(body, res)
		if err != nil {
			fmt.Println("unmarshal failed:", err)
		}
		startTheGame(res)
	}
	fmt.Println("Game Over")
	pressAnyButtonToContinue()
}

func browseRoom() {
	fmt.Println()

	URL := getUrl("/room") // 设置创建房间的URL
	// 构造请求的参数

	// 发送GET请求
	c, err := client.NewClient()
	if err != nil {
		fmt.Println("new client failed:", err)
	}

	req, resp := protocol.AcquireRequest(), protocol.AcquireResponse()
	req.SetRequestURI(URL)

	err = c.Do(context.Background(), req, resp)
	if err != nil {
		fmt.Println("do request failed:", err)
	}

	if resp.StatusCode() != http.StatusOK {
		fmt.Println("browse room failed: " + string(resp.Body()))
	} else {
		var res = &room.GetRoomListResp{}

		err = json.Unmarshal(resp.Body(), &res)
		if err != nil {
			fmt.Println("unmarshal failed:", err)
		}
		fmt.Printf("%-20s%-20s%-20s\n", "Name", "Capacity", "CurrentPlayer")
		for _, v := range res.Rooms {
			fmt.Printf("%-20s%-20d%-20d\n", v.Name, v.MaxPlayer, v.CurrentPlayer)
		}
	}
	pressAnyButtonToContinue()
}

func Show() {
	for {
		//清空屏幕
		fmt.Println("welcome to Tank War")
		fmt.Println("[1]Register [2]Login [3]CreateRoom [4]JoinRoom [5]BrowseRoom [6]Exit")
		char, _, _ := keyboard.GetSingleKey()
		switch char {
		case '1':
			register()
		case '2':

			login()
		case '3':
			createRoom()
		case '4':
			joinRoom()
		case '5':
			browseRoom()
		case '6':
			return
		default:
		}
		clearScreen()
	}
}

func getUrl(path string) string {
	return "http://" + consts.AgentServerAddress + ":" + strconv.Itoa(consts.AgentServerPort) + path
}

func pressAnyButtonToContinue() {
	fmt.Println("press any button to continue")
	_, _, err := keyboard.GetSingleKey()
	if err != nil {
		panic(err)
	}
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func startTheGame(r *room.JoinRoomResp) {

	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"tank_war"},
	}
	//connect to quic
	conn, err := quic.DialAddr(context.Background(), net.JoinHostPort(r.Address, strconv.Itoa(int(r.Port))), tlsConf, nil)
	if err != nil {
		log.Println("net work error", err)
		return
	}

	stream, err := conn.AcceptStream(context.Background())
	if err != nil {
		log.Println("accept stream error", err)
		return
	}

	//read client id
	var data = make([]byte, 1)
	_, err = stream.Read(data)
	if err != nil {
		log.Println(err)
	}

	//发送加入房间请求
	msg := &pb.JoinRoomReq{
		RoomId:     r.RoomID,
		PlayerId:   r.PlayerID,
		MaxPlayer:  r.MaxPlayer,
		RoomName:   r.RoomName,
		PlayerName: r.PlayerName,
	}

	data, err = proto.Marshal(msg)
	if err != nil {
		log.Println("proto marshal error", err)
		return
	}

	buffer := bytes.NewBuffer([]byte{})

	// write data length
	if err = binary.Write(buffer, binary.BigEndian, int32(len(data))); err != nil {
		log.Println("write data length error")
		return
	}
	// write data content
	if err = binary.Write(buffer, binary.BigEndian, data); err != nil {
		log.Println("write data content error", err)
		return
	}

	_, err = stream.Write(buffer.Bytes())

	if err != nil {
		log.Println("write err: ", err)
		return
	}

	screen, err := tcell.NewScreen()
	if err != nil {
		log.Println("new screen error", err)
		return
	}
	err = screen.Init()
	if err != nil {
		log.Println("init screen error", err)
		return
	}
	defer screen.Fini()

	c := handler.NewClient(stream, r.PlayerID, screen)

	screen.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorWhite))
	for {
		select {
		case <-c.Exit:
			screen.Fini()
			return
		default:
			screen.Clear()
			game.DrawBorder(screen)
			game.DrawRocks(screen)
			game.DrawPlayerList(screen)
			game.DrawTank(screen)
			game.DrawBullet(screen)
			game.DrawExplosion(screen)
			game.DrawBoard(screen)
			screen.Show()
			time.Sleep(17 * time.Millisecond)
		}
	}
}
