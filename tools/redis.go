package tools

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client
var RdbCtx context.Context

func InitRedisCli() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
	ctx := context.Background()
	RdbCtx = ctx
	Rdb = rdb
}
