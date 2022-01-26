package main

import (
	"encoding/json"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/gtcp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"golangProjects/frameworks/web/mvc/GoFrame/study/quick_start/网络服务开发/TCP组件/连接对象-消息包处理/自定义数据结构/types"
)

func main() {
	//数据上报客户端
	conn, err := gtcp.NewConn("127.0.0.1:8999")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	//使用JSON格式化数据字段
	info, err := json.Marshal(types.NodeInfo{
		Cpu:  float32(66.66),
		Host: "localhost",
		Ip: g.Map{
			"eth0": "192.168.1.100",
			"eht1": "114.114.10.11",
		},
		MemUsed:  15560320,
		MemTotal: 16333788,
		Time:     int(gtime.Timestamp()),
	})
	if err != nil {
		panic(err)
	}

	//使用 SendRecvPkg 发送消息包并接受返回
	if result, err := conn.SendRecvPkg(info); err != nil {
		if err.Error() == "EOF" {
			glog.Println("server closed")
		}
	} else {
		glog.Println(string(result))
	}
}
