package main

import (
	"log"
	"net"
	"net/rpc"
)

//Go语言的RPC框架有两个比较有特色的设计：一个是RPC数据打包时可以通过插件实现自定义的编码和
//解码；另一个是RPC建立在抽象的io.ReadWriteCloser接口之上的，我们可以将RPC架设在不同的通
//讯协议之上。这里我们将尝试通过官方自带的net/rpc/jsonrpc扩展实现一个跨语言的PPC。

type HelloService struct {
}

func (p *HelloService) Hello(requset string, reply *string) error {
	*reply = "server says:hello,client:" + requset
	return nil
}
func main() {
	rpc.RegisterName("HelloService", new(HelloService))

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
