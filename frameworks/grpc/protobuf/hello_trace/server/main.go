package main

import (
	"context"
	"fmt"
	"golang.org/x/net/trace"
	pb "golangProjects/frameworks/grpc/protobuf/hello_trace/proto/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
	"net/http"
)

const (
	Address = "127.0.0.1:50052"
)

// 定义并实现约定的接口
type HelloService struct {
}

// HelloService Hello服务
var helloService = HelloService{}

func (h HelloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := new(pb.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s", in.Name)

	return resp, nil
}

func startTrace() {
	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
		return true, true
	}
	go http.ListenAndServe(":50051", nil)
	grpclog.Println("Teace listen on 50051")
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("failed to listen : %v", err)
	}

	//实例化grpc server
	s := grpc.NewServer()

	//注册HelloService
	pb.RegisterHelloServer(s, helloService)

	//开启Trace
	go startTrace()

	grpclog.Println("Listen on " + Address)
	s.Serve(listen)
}
