package main

import (
	"fmt"
	"time"
)

func main() {
	//var wg sync.WaitGroup
	//wg.Add(2)
	/*	go func() {
			count(5,"sheep")
			wg.Done()
		}()
		go func() {
			go count(5,"bull")
			wg.Done()
		}()
		wg.Wait()
		fmt.Println(counter)*/
	c1 := make(chan string)
	c2 := make(chan string)
	c := make(chan string)
	go count(5, "sheep", c)
	/*for {
		message ,open  := <- c
		if !open {
			break
		}
	}*/
	/*for message := range c {
		fmt.Println(message)
	}*/
	go func() {
		for {
			c1 <- "sheep"
			time.Sleep(time.Millisecond * 500)
		}
	}()
	go func() {
		for {
			c2 <- "bull"
			time.Sleep(time.Millisecond * 2000)
		}
	}()
	for {
		select {
		case msg := <-c1:
			fmt.Println(msg)
		case msg := <-c2:
			fmt.Println(msg)
		}
	}
}

var counter int

func count(n int, animal string, c chan string) {
	/*	for i := 0 ; i <n ; i++ {
		//fmt.Println(i+1,animal)
		//counter += 1
	}*/
	for i := 0; i < n; i++ {
		c <- animal
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}
