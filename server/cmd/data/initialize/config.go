package initialize

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/spf13/viper"
	"tank_war/server/cmd/data/config"
	"tank_war/server/shared/consts"
)

func InitConfig() {
	v := viper.New()
	v.SetConfigFile(consts.DataConfigPath)
	if err := v.ReadInConfig(); err != nil {
		klog.Fatalf("read viper config failed: %s", err.Error())
	}
	if err := v.Unmarshal(&config.GlobalServerConfig); err != nil {
		klog.Fatalf("unmarshal err failed: %s", err.Error())
	}

}
