package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"log"
	"net"
	"strconv"
	"tank_war/server/cmd/user/config"
	"tank_war/server/cmd/user/pkg/md5"
	"tank_war/server/cmd/user/pkg/mysql"
	"tank_war/server/cmd/user/pkg/paseto"
	"tank_war/server/cmd/user/pkg/uuid"
	"tank_war/server/shared/kitex_gen/user/userservice"

	"tank_war/server/cmd/user/initialize"
)

func main() {
	initialize.InitLogger()
	initialize.InitConfig()
	IP, Port := initialize.InitFlag()
	r, info := initialize.InitRegistry(Port)
	db := initialize.InitDB()
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.GlobalServerConfig.Name),
		provider.WithExportEndpoint(config.GlobalServerConfig.OtelInfo.EndPoint),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	//TODO: 生成自己的pubkey和prikey
	tg, err := paseto.NewTokenGenerator()

	if err != nil {
		klog.Fatalf("new token generator failed: %s", err.Error())
	}

	svr := userservice.NewServer(&UserServiceImpl{
		EncryptManager: &md5.EncryptManager{Salt: config.GlobalServerConfig.MysqlInfo.Salt},
		MysqlManager:   mysql.NewUserManager(db, config.GlobalServerConfig.MysqlInfo.Salt),
		IDGenerator:    uuid.NewIDGenerator(),
		TokenGenerator: tg,
	},

		server.WithServiceAddr(utils.NewNetAddr("tcp", net.JoinHostPort(IP, strconv.Itoa(Port)))),
		server.WithRegistry(r),
		server.WithRegistryInfo(info),
		server.WithLimit(&limit.Option{MaxConnections: 2000, MaxQPS: 500}),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.GlobalServerConfig.Name}),
	)

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
