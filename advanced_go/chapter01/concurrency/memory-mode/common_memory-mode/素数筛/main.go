package main

import (
	"fmt"
)

// 返回生成自然数序列的管道：2,3,4,...
//函数内部启动一个Goroutine生产序列，返回对应的管道。
func GenerateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

//管道过滤器：删除能被素数整除的数
func PrimeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func main() {
	ch := GenerateNatural() // 自然数序列：2，3，4，...
	for i := 0; i < 100; i++ {
		prime := <-ch //新出现的素数
		fmt.Printf("%v : %v\n", i+1, prime)
		ch = PrimeFilter(ch, prime) //基于新素数构造的过滤器
	}
	//time.Sleep(1*time.Second)

}

/*

筛选法生成质数表（素数表）的基本思想如下：

假设有一个数组存放整数2 ~ N，

首先将2的倍数筛去（实际操作时可以将数组对应的值设置为0）

然后将3的倍数筛去，得到：

再一次将5的倍数筛去，7的倍数筛去，11的倍数筛去...........直到数组中所有的数都是质数。

 */