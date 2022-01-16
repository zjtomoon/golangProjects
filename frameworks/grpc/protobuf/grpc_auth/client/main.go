package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"time"

	pb "golangProjects/frameworks/grpc/protobuf/grpc_auth/proto/hello"
)

const (
	Address = "127.0.0.1:50052"

	//开启TLS认证
	OpenTLS = true
)

//customCredential 自定义认证
type customCredential struct {
}

//实现自定义认证接口
func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "101010",
		"appkey": "i am key",
	}, nil
}

// 自定义认证是否开启TLS
func (c customCredential) RequireTransportSecurity() bool {
	return OpenTLS
}

func interceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	grpclog.Printf("method=%s,req=%v,req=%v,duration=%s,error=%v\n", method, req, reply, time.Since(start), err)
	return err
}

func main() {
	////TLS连接
	//creds, err := credentials.NewClientTLSFromFile("../keys/server.pem", "server name")
	//if err != nil {
	//	grpclog.Fatalf("Failed to create TLS credentials %v", err)
	//}
	//conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(creds))
	//if err != nil {
	//	grpclog.Fatalln(err)
	//}
	//defer conn.Close()
	//
	////初始化客户端
	//c := pb.NewHelloClient(conn)
	//
	//req := &pb.HelloRequest{Name: "gRPC"}
	//
	//res, err := c.SayHello(context.Background(), req)
	//if err != nil {
	//	grpclog.Fatalln(err)
	//}
	//grpclog.Println(res.Message)
	//fmt.Println(res.Message)
	var err error
	var opts []grpc.DialOption

	if OpenTLS {
		//TLS连接
		creds, err := credentials.NewClientTLSFromFile("../keys/server.pem", "server name")
		if err != nil {
			grpclog.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	//指定自定义认证
	opts = append(opts, grpc.WithPerRPCCredentials(new(customCredential)))

	opts = append(opts, grpc.WithUnaryInterceptor(interceptor))

	conn, err := grpc.Dial(Address, opts...)
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()

	//初始化客户端
	c := pb.NewHelloClient(conn)

	//调用方法
	req := &pb.HelloRequest{Name: "gRPC"}
	res, err := c.SayHello(context.Background(), req)
	if err != nil {
		grpclog.Fatalln(err)
	}

	grpclog.Println(res.Message)
	fmt.Println(res.Message)
}
