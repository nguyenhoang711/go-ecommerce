package initialize

import (
	"context"
	"fmt"

	"github.com/devpenguin/go-ecommerce/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
    red := global.Config.Redis
    rdb := redis.NewClient(&redis.Options{
        Addr: fmt.Sprintf("%s:%v", red.Host, red.Port),
        Password: red.Password,
        DB: red.Database,
        PoolSize: red.PoolSize,
    })

    _, err := rdb.Ping(ctx).Result()
    if err != nil {
        global.Logger.Error("Redis initialization error: ", zap.Error(err))
    }
    global.Logger.Info("Initializing Redis Successfully")
    global.Rdb = rdb
}