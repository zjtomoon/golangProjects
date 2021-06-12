package main

import (
	"fmt"
	"time"
)

func do(){
	defer func(){
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	go da()
	go db()
	time.Sleep(3 * time.Second)
}
func da(){
	//panic("panic da")
	for i := 0 ; i < 10 ; i++ {
		fmt.Println(i)
	}
}


func db(){
	for i := 0 ; i < 10 ; i++ {
		fmt.Println(i)
	}
}
