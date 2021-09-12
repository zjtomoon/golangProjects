package main

import (
	"sync"
	"strconv"
	"fmt"
)
//sync.Map包 并发安全 开箱即用
//内置Store、Load、LoadStore、Delete、Range等操作方法
//对比 sync-mapBuildin
var m = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	for i := 0 ; i < 20 ; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key,n) //存值 必须使用sync.Map内置的方法设置键值对
			value,_ := m.Load(key) //读取 必须使用sync.Map提供的Load方法根据key取值
			fmt.Printf("k := %v,v := %v\n",key,value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
