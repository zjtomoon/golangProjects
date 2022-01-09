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

	_, err = c.Do("MSet", "name", "nick", "age", "18")
	if err != nil {
		fmt.Println("MSet error: ", err)
		return
	}

	r, err := redis.Strings(c.Do("MGet", "name", "age"))
	if err != nil {
		fmt.Println("MGet error: ", err)
		return
	}
	fmt.Println(r)
}
