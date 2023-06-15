package menu

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eiannone/keyboard"
	"github.com/gdamore/tcell/v2"
	"github.com/quic-go/quic-go"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"tank_war/client/game"
	"tank_war/client/handler"
	pb "tank_war/client/handler/pb/quic"
	"tank_war/server/cmd/api/biz/model/room"
	"tank_war/server/shared/kitex_gen/user"

	"tank_war/client/consts"
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

	var res = &user.LoginResp{}
	//unmarshal
	err = json.Unmarshal(body, res)
	if err != nil {
		fmt.Println("unmarshal failed:", err)
	}
	token = res.Token

	if status != http.StatusOK {
		fmt.Println("login failed " + string(body))
	} else {
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
	//req.AppendBody([]byte(fmt.Sprintf(`{"name":"%s","max_player":%d}`, roomName, capacity)))
	//req.SetBody([]byte(fmt.Sprintf(`{"name":"%s","max_player":%d}`, roomName, capacity)))

	// 发送POST请求
	err = c.DoTimeout(context.Background(), req, resp, time.Second*5)

	if err != nil {
		fmt.Println("do request failed:", err)
	}

	if resp.StatusCode() != http.StatusOK {
		fmt.Println("create room failed: " + string(resp.Body()))
	} else {
		fmt.Println(string(resp.Body()))
	}
	pressAnyButtonToContinue()
}

func joinRoom() {
	fmt.Println()

	var roomName string
	fmt.Println("please enter room name:")
	fmt.Scanln(&roomName)

	URL := getUrl("/room/join") // 设置创建房间的URL
	// 构造请求的参数
	ags := protocol.Args{}
	ags.Set("name", roomName)

	c, err := client.NewClient()
	if err != nil {
		fmt.Println("new client failed:", err)
	}

	// 发送POST请求
	status, body, err := c.Post(context.Background(), nil, URL, &ags)
	if err != nil {
		fmt.Println("do request failed:", err)
	}
	//TODO: get room info from agent server && start quic
	if status != http.StatusOK {
		fmt.Println("join room failed " + string(body))
	} else {
		fmt.Println(string(body))
	}
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
		fmt.Println("welcome to quic")
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

func startTheGame(host string, port int, roomName string, token string) {
	fmt.Println("start the quic...")
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	err = screen.Init()
	if err != nil {
		panic(err)
	}
	defer screen.Fini()
	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"tank_war"},
	}
	///connect to quic

	conn, err := quic.DialAddr(context.Background(), net.JoinHostPort(host, strconv.Itoa(port)), tlsConf, nil)
	if err != nil {
		panic(err)
	}

	stream, err := conn.AcceptStream(context.Background())
	if err != nil {
		panic(err)
	}

	////read client id
	//var data = make([]byte, 1)
	//_, err = stream.Read(data)
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//clientId := int32(data[0])
	////client write back to quic
	//_, err = stream.Write(data)
	//if err != nil {
	//	log.Println(err)
	//}

	msg := &pb.JoinRoomReq{
		RoomId: roomName,
		Token:  token,
	}

	data, err := json.Marshal(msg)
	if err != nil {
		klog.Fatalf("Marshal err: %v", err)
	}
	_, err = stream.Write(data)
	if err != nil {
		klog.Fatalf("write err: %v", err)
	}

	//read client id
	_, err = stream.Read(data)
	if err != nil {
		klog.Fatalf("read err: %v", err)
	}
	var res = &pb.JoinRoomResp{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		klog.Fatalf("unmarshal err: %v", err)
	}
	clientId := res.ClientId

	handler.NewClient(stream, clientId)

	screen.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorWhite))
	for {
		screen.Clear()
		game.DrawBorder(screen)
		game.DrawRocks(screen)
		game.DrawTank(screen)
		game.DrawBullet(screen)
		game.DrawExplosion(screen)
		game.DrawBoard(screen)
		screen.Show()
		time.Sleep(17 * time.Millisecond)
	}
}
