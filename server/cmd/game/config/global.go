package config

import "tank_war/server/shared/kitex_gen/user/userservice"

var (
	GlobalConsulConfig ConsulConfig
	GlobalServerConfig ServerConfig

	GlobalUserClient userservice.Client
)
