package main

import (
	"fmt"
	"github.com/gogf/gf/net/gtcp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"time"
)

func main() {
	//server
	go gtcp.NewServer("127.0.0.1:8999", func(conn *gtcp.Conn) {
		defer conn.Close()
		for {
			data, err := conn.Recv(-1)
			if len(data) > 0 {
				fmt.Println(string(data))
			}
			if err != nil {
				break
			}
		}
	}).Run()

	time.Sleep(time.Second)

	//client
	conn, err := gtcp.NewConn("127.0.0.1:8999")
	if err != nil {
		panic(err)
	}
	for i := 0; i < 10000; i++ {
		if err := conn.Send([]byte(gconv.String(i))); err != nil {
			glog.Error(err)
		}
		time.Sleep(time.Second)
	}
}
