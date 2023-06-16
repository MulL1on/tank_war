package initialize

import (
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
	"net"
	"strconv"
	"tank_war/server/cmd/api/config"
	"tank_war/server/shared/consts"
	"tank_war/server/shared/tools"
)

func InitConfig() {
	v := viper.New()
	v.SetConfigFile(consts.ApiConfigPath)
	if err := v.ReadInConfig(); err != nil {
		hlog.Fatalf("get config file failed,err", err)
	}

	if err := v.Unmarshal(&config.GlobalConsulConfig); err != nil {
		klog.Fatalf("unmarshal err failed", err.Error())
	}

	cfg := api.DefaultConfig()
	cfg.Address = net.JoinHostPort(
		config.GlobalConsulConfig.Host,
		strconv.Itoa(config.GlobalConsulConfig.Port),
	)

	client, err := api.NewClient(cfg)
	if err != nil {
		hlog.Fatalf("get consul client failed,err", err)
	}
	content, _, err := client.KV().Get(config.GlobalConsulConfig.Key, nil)
	if err != nil {
		hlog.Fatalf("get consul config failed,err", err)
	}

	err = sonic.Unmarshal(content.Value, &config.GlobalServerConfig)
	if config.GlobalServerConfig.Host == "" {
		config.GlobalServerConfig.Host, err = tools.GetLocalIPv4Address()
		if err != nil {
			hlog.Fatalf("get local ip failed,err", err)
		}
	}
	//hlog.Infof("config info", config.GlobalServerConfig.UserSrvInfo.Name)
}
