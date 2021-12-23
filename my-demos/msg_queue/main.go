package main

import (
	"fmt"
	"time"
)

type job struct {
	Index int
}

//生产者 消费者模式

func (t *job) doSomething() {
	time.Sleep(time.Second * 10)
	fmt.Println("complete work", t.Index)
}

var c1 chan job
var c2 chan chan job

func initJob() {
	for i := 0; i < 40; i++ {
		curJob[i] = job{}
		curJob[i].Index = i
		c1 <- curJob[i]
	}
}

func insertJob() {
	c := make(chan job, 1)
	select {
	case j := <-c1:
		c <- j
		c2 <- c
	}
}

func doWork(j job) {
	j.doSomething()
	insertJob()
}

func work() {
	for i := 0; i < 40; i++ {
		select {
		case c := <-c2:
			j := <-c
			go doWork(j)
		}
	}
}

var curJob []job

func main() {
	c1 = make(chan job, 40)
	c2 = make(chan chan job, 5)
	curJob = make([]job, 40)
	
	go initJob()
	insertJob()
	insertJob()
	insertJob()
	insertJob()
	insertJob()

	go initJob()
	work()
	time.Sleep(time.Second * 20)
	fmt.Println("ok")
}
