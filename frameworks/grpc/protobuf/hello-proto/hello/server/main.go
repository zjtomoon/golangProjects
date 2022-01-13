package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"

	pb "golangProjects/frameworks/grpc/protobuf/hello-proto/proto/hello"
)

const (
	// grpc服务地址
	Address = "127.0.0.1:50052"
)

type HelloService struct {
}

var helloService = HelloService{}

func (h HelloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := new(pb.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s.", in.Name)

	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}

	//实例化grpc server
	s := grpc.NewServer()

	//注册HelloService
	pb.RegisterHelloServer(s, helloService)

	grpclog.Infoln("Listen on " + Address)
	fmt.Println("Listen on", Address)

	s.Serve(listen)
}
