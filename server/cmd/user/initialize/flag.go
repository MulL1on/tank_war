package initialize

import (
	"flag"
	"github.com/cloudwego/kitex/pkg/klog"
	"tank_war/server/shared/consts"
	"tank_war/server/shared/tools"
)

func InitFlag() (string, int) {
	IP := flag.String(consts.IPFlagName, consts.IPFlagValue, consts.IPFlagUsage)
	Port := flag.Int(consts.PortFlagName, 0, consts.PortFlagUsage)

	//parsing flag, and if Port is 0, will automatically get an empty Port
	flag.Parse()
	if *Port == 0 {
		*Port, _ = tools.GetFreePort()
	}
	klog.Info("ip", *IP)
	klog.Info("port:", *Port)
	return *IP, *Port
}
