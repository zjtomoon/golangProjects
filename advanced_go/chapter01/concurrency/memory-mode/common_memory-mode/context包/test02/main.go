package main

import (
	"context"
	"fmt"
)


// 返回生成自然数序列的管道：2,3,4,...
func GenerateNatural(ctx context.Context) chan int {
	ch := make(chan int)
	go func() {
		for i := 2 ; ; i++ {
			select {
			case <- ctx.Done():
				return
			case ch <- i:
			}
		}
	}()
	return ch
}

//管道过滤器：删除能被素数整除的数
func PrimeFilter(ctx context.Context,in <- chan int,prime int) chan int {
	out := make(chan int)
	go func() {
		for  {
			if i := <- in ; i % prime != 0 {
				select {
				case <- ctx.Done():
					return
				case out <- i:
				}
			}
		}
	}()
	return out
}

func main() {
	// 通过Context控制后台goroutine状态
	ctx,cancel := context.WithCancel(context.Background())

	ch := GenerateNatural(ctx) //自然数序列：2，3，4，...
	for i := 0 ; i < 1000 ; i++ {
		prime := <- ch //新出现的素数
		ch = PrimeFilter(ctx,ch,prime) //基于新素数构造的过滤器
		fmt.Printf("%v : %v\n",i+1,prime)
	}
	cancel()
}
//当main函数完成工作前，通过调用cancel()来通知后台Goroutine退出，这样就避免了Goroutine的泄漏。