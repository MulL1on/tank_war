package main

import (
	"log"
	"os"
	"tank_war/client/menu"
)

var Token string

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	file, err := os.Create("./../tmp/client/log.txt")
	if err != nil {
		log.Fatal("无法创建日志文件:", err)
	}
	defer file.Close()

	// 设置日志输出位置为文件
	log.SetOutput(file)
	menu.Show()
}
