package main

import (
	"fmt"
	pb "golangProjects/advanced_go/chapter04/protobuf/hello-server/hello"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Dialing:", err)
	}

	var reply pb.String
	var request pb.String
	err = client.Call("HelloService.Hello", request, &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
