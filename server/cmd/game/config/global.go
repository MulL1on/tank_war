package config

import (
	"github.com/go-redis/redis/v8"
	"github.com/streadway/amqp"
	"tank_war/server/shared/kitex_gen/user/userservice"
)

var (
	GlobalConsulConfig ConsulConfig
	GlobalServerConfig ServerConfig

	GlobalUserClient userservice.Client
	Rdb              *redis.Client
	MqChan           *amqp.Channel
)
