package main

import (
	"fmt"
	"github.com/gogf/gf/net/gtcp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"time"
)

func main() {
	//Server
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
			return
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

//在这个示例中，Server每处理完毕一条请求之后便关闭链接。Client在第一条请求发送完毕后，由于连接池的IO复用特性，
// 下一次获取到的将是同一个连接对象，由于Server链接已主动关闭，第二次请求写入成功（其实并未成功发送到Server端，
// 需要通过下一次的读取操作才能检测到链接错误），但是读取却失败了（EOF表示目标连接关闭），
// 因此这个时候Client执行Close时将会销毁该连接操作对象，而不是进一步复用。
// 下一次再通过gtcp.NewPoolConn获得连接对象时，Client将会与Server创建一个新的连接进行数据通信。
// 所以你看到Client的端口一直在变化，那是因为该gtcp.Conn对象已经是一个新的连接对象，之前的连接对象已经被销毁。
//
//连接对象的IO复用涉及到十分微妙的连接状态变化问题，由于点对点网络通信本身是比较复杂的环境，
// 连接对象的状态随时可能被动发生着变化，因此，在使用gtcp连接池特性时，需要注意当通信错误产生时的连接对象重建机制，
// 一旦产生错误，立即丢弃（Close）该对象(gtcp.PoolConn)并重建（gtcp.NewPoolConn）。
