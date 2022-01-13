package main

import (
	"fmt"
	"golangProjects/frameworks/grpc/Hello/rpc_hello2/api"
	"log"
	"net/rpc"
)

type HelloService struct {
}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string

	err = client.Call(api.HelloServiceName+".Hello", "client says:hello,server", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
