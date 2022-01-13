package main

import (
	pb "golangProjects/advanced_go/chapter04/protobuf/hello-server/hello"
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (p *HelloService) Hello(request pb.String, reply *pb.String) error {
	reply.Value = "hello:" + request.GetValue()
	return nil
}

func main() {
	rpc.Register(new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTcp error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeConn(conn)
	}
}
