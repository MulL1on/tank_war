package config

import (
	"github.com/go-redis/redis/v8"
	"tank_war/server/shared/kitex_gen/user/userservice"
)

var (
	GlobalConsulConfig ConsulConfig
	GlobalServerConfig ServerConfig

	GlobalUserClient userservice.Client
	Rdb              *redis.Client
)
