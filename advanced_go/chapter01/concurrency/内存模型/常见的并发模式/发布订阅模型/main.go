package main

import (
	"time"
	"golangProjects/advanced_go/chapter01/concurrency/内存模型/常见的并发模式/发布订阅模型/pubsub"
)

func main() {
	p := pubsub.NewPublisher(100 * time.Millisecond,10)
	defer p.Close()

	all := p.Subscribe()
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s,ok := v.(string); ok {
			return strings.Contains(s,"golang")
		}
		return false
	})
}
