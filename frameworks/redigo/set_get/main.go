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

	_, err = c.Do("Set", "name", "nick")
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
