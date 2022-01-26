package main

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/gudp"
	"github.com/gogf/gf/os/gtime"
	"time"
)

func main() {
	//Server
	go gudp.NewServer("127.0.0.1:8999", func(conn *gudp.Conn) {
		defer conn.Close()
		for {
			data, err := conn.Recv(-1)
			if len(data) > 0 {
				if err := conn.Send(append([]byte("> "), data...)); err != nil {
					g.Log().Error(err)
				}
			}
			if err != nil {
				g.Log().Error(err)
			}
		}
	}).Run()

	time.Sleep(time.Second)

	//Client
	for {
		if conn, err := gudp.NewConn("127.0.0.1:8999"); err == nil {
			if b, err := conn.SendRecv([]byte(gtime.Datetime()), -1); err == nil {
				fmt.Println(string(b), conn.LocalAddr(), conn.RemoteAddr())
			} else {
				g.Log().Error(err)
			}
			conn.Close()
		} else {
			g.Log().Error(err)
		}
		time.Sleep(time.Second)
	}
}
