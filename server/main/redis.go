package main

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func initPool(address string, maxIdle int, maxActive int, IdleTimeout time.Duration) {
	pool = &redis.Pool{
		MaxIdle: maxIdle,
		MaxActive: maxActive,
		IdleTimeout: IdleTimeout, // 最大空闲时间
		Dial: func()(redis.Conn, error) {
			// 初始化链接的代码， 链接哪个 ip 的 redis
			return redis.Dial("tcp", address)
		},
	}
}