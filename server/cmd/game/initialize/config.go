package initialize

import (
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
	"net"
	"strconv"
	"tank_war/server/cmd/game/config"
	"tank_war/server/shared/consts"
	"tank_war/server/shared/tools"
)

func InitConfig() {
	v := viper.New()
	v.SetConfigFile(consts.GameConfigPath)
	if err := v.ReadInConfig(); err != nil {
		klog.Fatalf("read viper config failed: %s", err.Error())
	}
	if err := v.Unmarshal(&config.GlobalConsulConfig); err != nil {
		klog.Fatalf("unmarshal err failed: %s", err.Error())
	}
	klog.Infof("Config Info: %v", config.GlobalConsulConfig)

	cfg := api.DefaultConfig()
	cfg.Address = net.JoinHostPort(
		config.GlobalConsulConfig.Host,
		strconv.Itoa(config.GlobalConsulConfig.Port))
	consulClient, err := api.NewClient(cfg)
	if err != nil {
		klog.Fatalf("new consul client failed: %s", err.Error())
	}
	content, _, err := consulClient.KV().Get(config.GlobalConsulConfig.Key, nil)
	if err != nil {
		klog.Fatalf("consul kv failed: %v", err.Error())
	}
	err = sonic.Unmarshal(content.Value, &config.GlobalServerConfig)
	if err != nil {
		klog.Fatalf("sonic unmarshal config failed: %s", err.Error())
	}

	if config.GlobalServerConfig.Host == "" {
		config.GlobalServerConfig.Host, err = tools.GetLocalIPv4Address()
		if err != nil {
			klog.Fatalf("get localIPv4Addr failed: %s", err.Error())
		}
	}
}
