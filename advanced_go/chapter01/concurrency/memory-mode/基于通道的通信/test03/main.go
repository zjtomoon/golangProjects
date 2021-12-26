package main

//var done = make(chan bool,1) //无法打印
var done = make(chan bool)
var msg string

func aGoroutine()  {
	msg = "hello,world"
	<- done
}

func main() {
	go aGoroutine()
	done <- true
	//对于从无缓冲Channel进行的接收，发生在对该Channel进行的发送完成之前
	//交换两个goroutine的接收和发送操作也是可以的，但是很危险
	//如果该Channel为带缓冲的（done = make(chan bool,1)），main线程的done<-true接收操作将不会被后台的<-done接收，操作阻塞
	println(msg)
}
