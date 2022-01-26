package main

import (
	"fmt"
	"github.com/gogf/gf/net/gtcp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"time"
)

//gtcp模块提供了连接池的特性，由gtcp.PoolConn对象实现，连接池缓存固定存活时间为600秒，
// 且内部实现了数据发送时的断开重连机制。
// 连接池非常适合于频繁的短链接操作且连接并发量大的场景。我们接下来使用两个示例来演示一下连接池的作用。
func main() {
	//server
	go gtcp.NewServer("127.0.0.1:8999", func(conn *gtcp.Conn) {
		defer conn.Close()
		for {
			data, err := conn.Recv(-1)
			if len(data) > 0 {
				if err := conn.Send(append([]byte("> "), data...)); err != nil {
					fmt.Println(err)
				}
			}
			if err != nil {
				break
			}
		}
	}).Run()

	time.Sleep(time.Second)

	//Client
	for {
		if conn, err := gtcp.NewPoolConn("127.0.0.1:8999"); err == nil {
			if b, err := conn.SendRecv([]byte(gtime.Datetime()), -1); err == nil {
				fmt.Println(string(b), conn.LocalAddr(), conn.RemoteAddr())
			} else {
				fmt.Println(err)
			}
			conn.Close()
		} else {
			glog.Error(err)
		}
		time.Sleep(time.Second)
	}
}

//可以看到，Client的端口一直未变，每一次通过gtcp.NewConn("127.0.0.1:8999")获得的都是同一个gtcp.Conn对象，
// 且每一次conn.Close()时并不是真正的关闭连接，而是将该对象重新丢回到连接池里循环使用。
