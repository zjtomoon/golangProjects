package main

var done = make(chan bool)
var msg string

func aGoroutine()  {
	msg = "hello,world"
	//done <- true
	<-done
}

func main() {
	go aGoroutine()
	//<- done
	done <- true
	println(msg)
}
