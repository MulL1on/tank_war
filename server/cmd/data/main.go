package main

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	"tank_war/server/cmd/data/config"
	"tank_war/server/cmd/data/initialize"
	"tank_war/server/cmd/data/pkg/mysql"
	"tank_war/server/shared/consts"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	file, err := os.Create("./log.txt")
	if err != nil {
		log.Fatal("无法创建日志文件:", err)
	}
	defer file.Close()

	// 设置日志输出位置为文件
	log.SetOutput(file)

	initialize.InitConfig()
	db := initialize.InitDB()
	m := mysql.NewUserManager(db)
	cfg := config.GlobalServerConfig.RabbitMQInfo
	conn, err := amqp.Dial(fmt.Sprintf(consts.RabbitMqUrl, cfg.Username, cfg.Password, cfg.Host, cfg.Port))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Println("connected to rabbitmq server successfully")

	// 创建一个通道
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	// 声明一个队列
	queue, err := ch.QueueDeclare(
		"user_data",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	for d := range msgs {
		var u = &mysql.User{}
		err := json.Unmarshal(d.Body, u)
		if err != nil {
			log.Println(err)
			continue
		}
		m.UpdateUser(u)
	}
}
