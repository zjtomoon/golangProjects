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

	//设置过期时间
	_, err = c.Do("expire", "names", 5)
	if err != nil {
		fmt.Println("expore error", err)
		return
	}

	fmt.Println("expired")

}
