package main

import (
	"sync"
	"testing"
)

//sync.Pool的回收是有的，它是在系统自动GC的时候，触发pool.go中的poolCleanup函数
//sync.Pool其实不适合用来做持久保存的对象池（比如连接池）。
// 它更适合用来做临时对象池，目的是为了降低GC的压力。

var bytePool = sync.Pool{
	New: newPool,
}

func newPool() interface{} {
	b := make([]byte, 1024)
	return &b
}

func BenchmarkAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		obj := make([]byte, 1024)
		_ = obj
	}
}

func BenchmarkPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		obj := bytePool.Get().(*[]byte)
		_ = obj
		bytePool.Put(obj)
	}
}
