package initialize

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/go-redis/redis/v8"
	"tank_war/server/cmd/game/config"
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
		klog.Fatalf("connect to redis failed,err:%v", err)
	}
	config.Rdb = rdb
}
