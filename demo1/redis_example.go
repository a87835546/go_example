package demo1

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisExample struct {
}

var ctx = context.Background()
var rdb = redis.Client{}

func init() {
	rdb = *redis.NewClient(&redis.Options{Addr: "192.168.0.229:6379", Password: "", DB: 1})
	res, err := rdb.Ping(ctx).Result()
	err = rdb.Set(ctx, "name", "123", 0).Err()
	if err != nil {
		fmt.Println("ping 出错：", err)
	}
	fmt.Println(res)
}
func SetValue(key string, value any) {
	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		fmt.Println("set 出错：", err)
	}
}
func HSetValue(key string, value any) {
	err := rdb.HSet(ctx, key, value).Err()
	if err != nil {
		fmt.Println("set hash 出错：", err)
	}
}
func SetValueWithEx(key, value string, ex int64) {
	err := rdb.Set(ctx, key, value, time.Duration(ex)).Err()
	if err != nil {
		fmt.Println("set 出错：", err)
	}
}
func Get(key string) string {
	res, err := rdb.Get(ctx, key).Result()
	if err != nil {
		fmt.Println("get 出错：", err)
	} else {
		fmt.Printf("get res --->>> %v\n\n", res)
	}
	return res
}
