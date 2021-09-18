package main

import (
	"math/rand"
	"time"
	"strconv"
	"fmt"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	const Max = 100000
	const NumReceivers = 10
	const NumSenders = 1000

	dataCh := make(chan int,100)
	stopCh := make(chan struct{})

	//必须是1个缓存通道
	toStop := make(chan string,1)

	var stoppedBy string

	//moderator
	go func() {
		stoppedBy = <- toStop
		close(stopCh)
	}()

	//senders
	for i := 0 ; i < NumReceivers ; i++ {
		go func(id string) {
			for {
				value := rand.Intn(Max)
				if value == 0 {
					select {
					case toStop <- "sender#" + id:
					default:
					}
					return
				}
				select {
				case <- stopCh:
					return
				case dataCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}

	//receivers
	for i := 0 ; i < NumReceivers ; i++ {
		go func(id string) {
			for {
				select {
				case <- stopCh:
					return
				case value := <- dataCh:
					if value == Max -1 {
						select {
						case toStop <- "receive#" + id:
						default:
						}
						return
					}
					fmt.Println(value)
				}
			}
		}(strconv.Itoa(i))
	}

	select {
	case <- time.After(time.Hour):
		
	}
}
