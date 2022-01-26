package main

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/gtcp"
	"github.com/gogf/gf/util/gconv"
	"time"
)

func main() {
	address := "127.0.0.1:8999"
	crtFile := "server.crt"
	keyFile := "server.key"

	//TLS server
	go gtcp.NewServerKeyCrt(address, crtFile, keyFile, func(conn *gtcp.Conn) {
		defer conn.Close()
		for {
			data, err := conn.Recv(-1)
			if len(data) > 0 {
				fmt.Println(string(data))
			}
			if err != nil {
				g.Log().Error(err)
				break
			}
		}
	}).Run()
	time.Sleep(time.Second)

	//Client
	tlsConfig, err := gtcp.LoadKeyCrt(crtFile, keyFile)
	if err != nil {
		panic(err)
	}
	//忽略证书校验
	tlsConfig.InsecureSkipVerify = true

	conn, err := gtcp.NewConnTLS(address, tlsConfig)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	for i := 0; i < 10; i++ {
		if err := conn.Send([]byte(gconv.String(i))); err != nil {
			g.Log().Error(err)
			//glog.Error()
		}

		time.Sleep(time.Second)
		if i == 5 {
			conn.Close()
			break
		}
	}
	time.Sleep(5 * time.Second)
}
