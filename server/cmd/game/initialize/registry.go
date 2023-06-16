package initialize

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hashicorp/consul/api"
	"net"
	"strconv"
	"tank_war/server/cmd/game/config"
	"tank_war/server/shared/consts"
	"time"
)

func InitRegistry() {
	// TODO: InitRegistry
	cfg := api.DefaultConfig()
	cfg.Address = net.JoinHostPort(
		config.GlobalConsulConfig.Host,
		strconv.Itoa(config.GlobalConsulConfig.Port))
	//helth check
	check := &api.AgentServiceCheck{
		TTL:                            consts.ConsulCheckTTL,
		DeregisterCriticalServiceAfter: "1s",
		Timeout:                        consts.ConsulCheckTimeout,
	}

	client, err := api.NewClient(cfg)
	if err != nil {
		klog.Fatalf("new consul register failed: %s", err.Error())
	}

	r := &api.AgentServiceRegistration{
		ID:      "game_srv:" + config.GlobalServerConfig.Host + ":" + strconv.Itoa(config.GlobalServerConfig.Port),
		Address: config.GlobalServerConfig.Host,
		Name:    "game_srv",
		Port:    config.GlobalServerConfig.Port,
		Check:   check,
		Meta:    map[string]string{"status": "free"},
	}

	if err := client.Agent().ServiceRegister(r); err != nil {
		klog.Fatalf("register service failed: ", err)
	}
	go handleTTLCheck(r)
}

func handleTTLCheck(r *api.AgentServiceRegistration) {
	// 获取Consul客户端
	cfg := api.DefaultConfig()
	cfg.Address = net.JoinHostPort(
		config.GlobalConsulConfig.Host,
		strconv.Itoa(config.GlobalConsulConfig.Port))
	client, err := api.NewClient(cfg)
	if err != nil {
		klog.Fatalf("new consul client failed: %s", err.Error())
	}

	// 定期发送TTL检查
	ticker := time.NewTicker((15 * time.Second * 9) / 10) // 以5秒为间隔发送TTL检查
	for range ticker.C {
		// 通过Consul API更新TTL检查
		err := client.Agent().PassTTL("service:"+r.ID, "TTL check passed")
		if err != nil {
			klog.Errorf("TTL check failed: %s", err.Error())
		}
	}
}
