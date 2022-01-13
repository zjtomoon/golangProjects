package main

import (
	"context"
	"fmt"
	pb "golangProjects/frameworks/grpc/protobuf/grpc_auth/proto/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

const (
	Address = "127.0.0.1:50052"
)

func main() {
	//TLS连接
	creds, err := credentials.NewClientTLSFromFile("../keys/server.pem", "server name")
	if err != nil {
		grpclog.Fatalf("Failed to create TLS credentials %v", err)
	}
	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(creds))
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()

	//初始化客户端
	c := pb.NewHelloClient(conn)

	req := &pb.HelloRequest{Name: "gRPC"}

	res, err := c.SayHello(context.Background(), req)
	if err != nil {
		grpclog.Fatalln(err)
	}
	grpclog.Println(res.Message)
	fmt.Println(res.Message)
}
