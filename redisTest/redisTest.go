package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

func main() {
	// Redis 配置
	redisConfig := struct {
		Host     string
		Port     int
		Password string
		DB       int
		PoolSize int
	}{
		Host:     "127.0.0.1",
		Port:     6379,
		Password: "123456", // 如果没有密码，则设置为空字符串 ""
		DB:       0,
		PoolSize: 10,
	}

	addr := fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port)
	redisConn := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
		PoolSize: redisConfig.PoolSize,
	})

	// 测试连接
	res, err := redisConn.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to ping redis, err: %v", err)
	}
	fmt.Printf("Ping response: %s\n", res)
}
