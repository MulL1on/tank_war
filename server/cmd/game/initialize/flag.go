package initialize

import (
	"flag"
	"github.com/cloudwego/kitex/pkg/klog"
	"tank_war/server/cmd/game/config"
	"tank_war/server/shared/consts"
	"tank_war/server/shared/tools"
)

func InitFlag() {
	Port := flag.Int(consts.PortFlagName, 0, consts.PortFlagUsage)

	//parsing flag, and if Port is 0, will automatically get an empty Port
	flag.Parse()
	if *Port == 0 {
		*Port, _ = tools.GetFreePortInRange(config.GlobalServerConfig.Host)
	}
	klog.Info("port:", *Port)
	config.GlobalServerConfig.Port = *Port
}
