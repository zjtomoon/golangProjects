package main

import (
	"github.com/gogf/gf/net/gtcp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtimer"
	"golangProjects/frameworks/web/mvc/GoFrame/study/quick_start/网络服务开发/TCP组件/连接对象-通信开发进阶/funcs"
	"golangProjects/frameworks/web/mvc/GoFrame/study/quick_start/网络服务开发/TCP组件/连接对象-通信开发进阶/types"
	"time"
)

func main() {
	gtcp.NewServer("127.0.0.1:8999", func(conn *gtcp.Conn) {
		defer conn.Close()

		//测试消息，10秒后让客户端主动退出
		gtimer.SetTimeout(10*time.Second, func() {
			funcs.SendPkg(conn, "doexit")
		})
		for {
			msg, err := funcs.RecvPkg(conn)
			if err != nil {
				if err.Error() == "EOF" {
					glog.Println("client closed")
				}
				break
			}
			switch msg.Act {
			case "hello":
				onClientHello(conn, msg)
			case "heartbeat":
				onClietnHeartBeat(conn, msg)
			default:
				glog.Errorf("invalid message : %v", msg)
				break
			}
		}
	}).Run()
}

func onClientHello(conn *gtcp.Conn, msg *types.Msg) {
	glog.Printf("hello message from [%s]: %s", conn.RemoteAddr().String(), msg.Data)
	funcs.SendPkg(conn, msg.Act, "Nice to meet you!")
}

func onClietnHeartBeat(conn *gtcp.Conn, msg *types.Msg) {
	glog.Printf("heartbeat from [%s]", conn.RemoteAddr().String())
}
