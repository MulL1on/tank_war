package initialize

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-redis/redis/v8"
	"tank_war/server/cmd/api/config"
	"time"
)

func InitRdb() {
	cfg := config.GlobalServerConfig.RedisInfo
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Username: cfg.Username,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		hlog.Fatalf("connect to redis failed,err:%v", err)
	}
	config.Rdb = rdb
	hlog.Info("initialize redis successfully.")
}
