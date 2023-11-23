package tools

import (
	"fmt"
	"testing"
)

func TestInitRedisCli(t *testing.T) {
	InitConfig()
	InitRedisCli()
	testKey := "testgo"
	err := Rdb.Set(RdbCtx, testKey, "testgov", -1).Err()
	if err != nil {
		panic(err)
	} else {
		if v, err := Rdb.Get(RdbCtx, testKey).Result(); err == nil {
			fmt.Println(fmt.Sprintf("get the key value: %s", v))
		} else {
			panic(err)
		}
	}
}
