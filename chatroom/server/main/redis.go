package main

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func initPool(address string, maxIdle int, maxActive int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return (redis.Dial("tcp", address))
		},
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeout,
	}
}
