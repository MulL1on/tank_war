package initialize

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/streadway/amqp"
	"tank_war/server/cmd/game/config"
	"tank_war/server/shared/consts"
)

func InitMq() *amqp.Connection {
	cfg := config.GlobalServerConfig.RabbitMQInfo
	conn, err := amqp.Dial(fmt.Sprintf(consts.RabbitMqUrl, cfg.Username, cfg.Password, cfg.Host, cfg.Port))
	klog.Infof("connect to rabbitmq url:", fmt.Sprintf(consts.RabbitMqUrl, cfg.Username, cfg.Password, cfg.Host, cfg.Port))
	if err != nil {
		klog.Fatal(err)
	}
	return conn

}
