package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

//连接池例子
var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     16,
		MaxActive:   1024,
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	c := pool.Get()
	defer c.Close()

	_, err := c.Do("Set", "name", "nick2")
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.String(c.Do("Get", "name"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r)
}
