package tools

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
)

var Rdb *redis.Client
var RdbCtx context.Context

func InitRedisCli() {
	redisPort, _ := strconv.Atoi(Config["REDIS_DB"])
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", Config["REDIS_HOST"], Config["REDIS_PORT"]),
		Password: Config["REDIS_PASSWORD"], // 没有密码，默认值
		DB:       redisPort,                // 默认DB 0
	})
	ctx := context.Background()
	RdbCtx = ctx
	Rdb = rdb
}
