package initialize

import (
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"
	"net"
	"strconv"
	"tank_war/server/cmd/api/config"
	"tank_war/server/shared/consts"

	"github.com/hashicorp/consul/api"
)

func InitRegistry(Port int) {
	// TODO: InitRegistry
	cfg := api.DefaultConfig()
	cfg.Address = net.JoinHostPort(
		config.GlobalConsulConfig.Host,
		strconv.Itoa(config.GlobalConsulConfig.Port))
	check := &api.AgentServiceCheck{
		Interval:                       consts.ConsulCheckInterval,
		Timeout:                        consts.ConsulCheckTimeout,
		DeregisterCriticalServiceAfter: consts.ConsulCheckDeregisterCriticalServiceAfter,
	}

	client, err := api.NewClient(cfg)
	if err != nil {
		klog.Fatalf("new consul register failed: %s", err.Error())
	}
	sf, err := snowflake.NewNode(3)

	r := &api.AgentServiceRegistration{
		ID:      sf.Generate().Base36(),
		Address: config.GlobalServerConfig.Host,
		Name:    "game_srv",
		Port:    Port,
		Check:   check,
	}

	if err := client.Agent().ServiceRegister(r); err != nil {
		panic(err)
	}
	klog.Infof("register service success")
}
