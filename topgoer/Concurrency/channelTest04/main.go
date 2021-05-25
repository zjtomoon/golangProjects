package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	//开启goroutinue将0~100的数发送到ch1中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	//开启goroutinue从ch1接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			j, ok := <-ch1 //关闭通道后再取值，ok=false
			if !ok {
				break
			}
			ch2 <- j * j
		}
		close(ch2)
	}()

	//在主goroutinue中从ch2中接收值并打印
	for k := range ch2 {
		fmt.Println(k)
	}
}