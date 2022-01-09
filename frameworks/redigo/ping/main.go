package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn failed,err: ", err)
		return
	}
	defer c.Close()
	r, _ := redis.String(c.Do("ping"))
	fmt.Println(r)
}
