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

	//队列
	_, err = c.Do("lpush", "Queue", "nick", "dawn", 9)
	if err != nil {
		fmt.Println("lpush error: ", err)
		return
	}
	for {
		r, err := redis.String(c.Do("lpop", "Queue"))
		if err != nil {
			fmt.Println("lpop error : ", err)
			break
		}
		fmt.Println(r)
	}
	r2, err := redis.Int(c.Do("llen", "Queue"))
	if err != nil {
		fmt.Println("llen error:", err)
		return
	}
	fmt.Println(r2)
}
