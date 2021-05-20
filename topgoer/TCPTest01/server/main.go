package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn){
	defer conn.Close()
	reader := bufio.NewReader(conn)
	var buf[1024]byte
	for {
		n,err := reader.Read(buf[:])
		if err != nil{
			fmt.Println("read from client failed,err:",err)
			break
		}
		if err == io.EOF {
			break
		}
		recvStr :=string(buf[:n])
		fmt.Println("收到client发来的消息:",recvStr)
	}
}

func main(){
	listen,err := net.Listen("tcp","127.0.0.1:30000")
	if err != nil{
		fmt.Println("listen failed,err :",err)
		return
	}
	defer listen.Close()
	for {
		conn,err := listen.Accept()
		if err != nil{
			fmt.Println("accept failed,err:",err)
			continue
		}
		go process(conn)
	}
}
