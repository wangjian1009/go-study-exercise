package main

import (
	"fmt"
	"strconv"

	"github.com/garyburd/redigo/redis"
)

// use redis to save key-value(1W-50W)
func main() {
	redisPool := &redis.Pool{
		MaxIdle:     10,
		MaxActive:   1000,
		IdleTimeout: 60,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
	defer redisPool.Close()

	reConn := redisPool.Get()
	defer reConn.Close()

	// init 10000 size "1"
	var value = ""
	for j := 0; j < 100000; {
		value += "1"
		j++
	}

	for i := 0; i < 100000; i++ {
		_, err := reConn.Do("set", "key"+strconv.Itoa(i+1), value)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
