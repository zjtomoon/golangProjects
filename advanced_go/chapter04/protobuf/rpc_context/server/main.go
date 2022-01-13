package main

import (
	"log"
	"net"
	"net/rpc"
)

//上下文信息

//基于上下文我们可以针对不同客户端提供定制化的RPC服务。我们可以通过为每个链接提供独立的RPC
//服务来实现对上下文特性的支持。

type HelloService struct {
	conn net.Conn
}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request + ",from" + p.conn.RemoteAddr().String()
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	//为每个链接启动独立的RPC服务
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go func() {
			defer conn.Close()
			p := rpc.NewServer()
			p.Register(&HelloService{conn: conn})
			p.ServeConn(conn)
		}()
	}
}
