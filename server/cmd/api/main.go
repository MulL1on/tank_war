// Code generated by hertz generator.

package main

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	cfg "github.com/hertz-contrib/http2/config"
	"github.com/hertz-contrib/http2/factory"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/hertz-contrib/pprof"
	"tank_war/server/cmd/api/config"
	"tank_war/server/cmd/api/initialize"
	"tank_war/server/cmd/api/initialize/rpc"
	"time"
)

func main() {
	initialize.InitLogger()
	initialize.InitConfig()
	r, info := initialize.InitRegistry()
	tracer, trcCfg := hertztracing.NewServerTracer()
	initialize.InitRdb()
	rpc.Init()

	h := server.New(
		tracer,
		server.WithALPN(true),
		server.WithHostPorts(fmt.Sprintf(":%d", config.GlobalServerConfig.Port)),
		server.WithRegistry(r, info),
		server.WithHandleMethodNotAllowed(true),
	)

	h.AddProtocol("h2", factory.NewServerFactory(
		cfg.WithReadTimeout(time.Minute),
		cfg.WithDisableKeepAlive(false)))

	pprof.Register(h)
	h.Use(hertztracing.ServerMiddleware(trcCfg))

	register(h)
	h.Spin()
}
