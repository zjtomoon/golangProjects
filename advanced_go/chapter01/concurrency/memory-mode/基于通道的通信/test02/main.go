package main

var done = make(chan bool)
var msg string

func aGoroutine()  {
	msg = "你好，世界"
	close(done)
}

func main() {
	go aGoroutine()
	<- done
	println(msg)
}
/*
	若在关闭通道后继续从中接收数据，接收者就会收到该通道返回的零值。
	用close(done)关闭通道代替done <- true依然能保证该程序产生相同的行为。

	对于从无缓存通道进行的接收，发生在对该通道进行的发送完成之前。
 */