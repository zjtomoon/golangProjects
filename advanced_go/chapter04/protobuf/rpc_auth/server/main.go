package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {
	conn    net.Conn
	isLogin bool
}

func ServeHelloService(conn net.Conn) {
	defer conn.Close()
	p := rpc.NewServer()
	p.Register(&HelloService{conn: conn})
	p.ServeConn(conn)
}

//为RPC服务增加简单的登陆状态的验证：
//这样可以要求在客户端链接RPC服务时，首先要执行登陆操作，登陆成功后才能正常执行其他的服务。
func (p *HelloService) Login(request string, reply *string) error {
	if request != "user:password" {
		log.Println("login failed")
		return fmt.Errorf("auth failed")
	}
	log.Println("login ok")
	p.isLogin = true
	return nil
}

func (p *HelloService) Hello(request string, reply *string) error {
	if !p.isLogin {
		return fmt.Errorf("please login")
	}
	*reply = "hello:" + request + ", from " + p.conn.RemoteAddr().String()
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go ServeHelloService(conn)
	}
}
