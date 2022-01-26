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
	conn, err := gtcp.NewConn("127.0.0.1:8999")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	//心跳信息
	gtimer.SetInterval(time.Second, func() {
		if err := funcs.SendPkg(conn, "heartbeat"); err != nil {
			panic(err)
		}
	})

	//测试消息，3秒后服务端发送hello消息
	gtimer.SetTimeout(3*time.Second, func() {
		if err := funcs.SendPkg(conn, "hello", "My name's John"); err != nil {
			panic(err)
		}
	})
	for {
		msg, err := funcs.RecvPkg(conn)
		if err != nil {
			if err.Error() == "EOF" {
				glog.Println("server closed")
			}
			break
		}
		switch msg.Act {
		case "hello":
			onServerHello(conn, msg)
		case "doexit":
			onServerDoExit(conn, msg)
		case "heartbeat":
			onServerHeartBeat(conn, msg)
		default:
			glog.Errorf("invalid message: %v", msg)
			break
		}
	}
}

func onServerHello(conn *gtcp.Conn, msg *types.Msg) {
	glog.Printf("hello response message from [%s]: %s", conn.RemoteAddr().String(), msg.Data)
}

func onServerHeartBeat(conn *gtcp.Conn, msg *types.Msg) {
	glog.Printf("hearbeat from [%s]", conn.RemoteAddr().String())
}

func onServerDoExit(conn *gtcp.Conn, msg *types.Msg) {
	glog.Printf("exit command from [%s]", conn.RemoteAddr().String())
	conn.Close()
}
