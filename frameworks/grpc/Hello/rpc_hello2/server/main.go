package main

import (
	"golangProjects/frameworks/grpc/Hello/rpc_hello2/api"
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (p *HelloService) Hello(requset string, reply *string) error {
	*reply = "server says:hello,client:" + requset
	return nil
}

func main() {
	api.RegisterHelloService(new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeConn(conn)
	}
}
