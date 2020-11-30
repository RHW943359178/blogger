package redis

import (
	"github.com/go-redis/redis"
	"log"
)

var (
	rdb *redis.Client
)

//	初始化连接
func InitClient(addr string) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
		PoolSize: 100,
	})

	//	连接验证
	_, err = rdb.Ping().Result()
	if err != nil {
		return
	}
	log.Println("redis 连接成功！")
	return
}
