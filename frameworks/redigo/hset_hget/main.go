package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn failed ,err: ", err)
		return
	}
	defer c.Close()

	_, err = c.Do("HSet", "names", "nick", "suoning")
	if err != nil {
		fmt.Println("hset error", err)
		return
	}
	r, err := redis.String(c.Do("HGet", "names", "nick"))
	if err != nil {
		fmt.Println("hget error: ", err)
		return
	}
	fmt.Println(r)
}
