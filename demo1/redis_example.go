package demo1

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

type RedisExample struct {
}

var ctx = context.Background()
var rdb = &redis.Client{}

func init() {
	log.Println("start ... redis")
	rdb = redis.NewClient(&redis.Options{Addr: "13.215.251.145:6379", Password: "", DB: 0})
	res, err := rdb.Ping(ctx).Result()
	err = rdb.Set(ctx, "name", "123", 0).Err()
	if err != nil {
		log.Println("ping 出错：", err)
	}
	log.Println("redis.... ", res)
}
func SetValue(key string, value any) {
	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		log.Println("set 出错：", err)
	}
}
func HSetValue(key string, value any) {
	err := rdb.HSet(ctx, key, value).Err()
	if err != nil {
		log.Println("set hash 出错：", err)
	}
}
func SetValueWithEx(key, value string, ex int64) {
	err := rdb.Set(ctx, key, value, time.Duration(ex)).Err()
	if err != nil {
		log.Println("set 出错：", err)
	}
}
func Get(key string) string {
	res, err := rdb.Get(ctx, key).Result()
	if err != nil {
		log.Println("get 出错：", err)
	} else {
		log.Printf("get res --->>> %v\n\n", res)
	}
	return res
}
