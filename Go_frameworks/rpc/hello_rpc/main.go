package main

import (
	"flag"
	"github.com/smallnest/rpcx/server"

	example "github.com/rpcx-ecosystem/rpcx-examples3"
)

var (
	addr = flag.String("addr","localhost:8080","server address")
)

func main() {
	flag.Parse()

	s := server.NewServer()
	s.RegisterName("Arith",new(example.Arith),"")
	s.Register(new(example.Arith),"")
	s.Serve("tcp",*addr)
}
