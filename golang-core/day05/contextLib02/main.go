package main

import (
	"context"
	"fmt"
	"time"
)

type otherContext struct {
	context.Context
}

func main() {
	//使用context.Backgrond()构建一个WithCancel类型的上下文
	ctxa,cancel := context.WithCancel(context.Background())

	/*
		ctxa 内部状态：
		ctxa = & {
		Context: new(emptyCtx)
		}
	 */

	//work模拟运行并检测前端的退出通知
	go work(ctxa,"work1")

	//使用WithDeadline包装前面的上下文对象ctxa
	tm := time.Now().Add(3 * time.Second)
	ctxb,_ := context.WithDeadline(ctxa,tm)

	/*
		ctxb内部状态：
		ctxb = &timerCtx {
		cancelCtx:ctxa
		deadline:tm
		}
		同时触发ctxa,在children中维护ctxb作为子节点
	 */

	go work(ctxb,"work2")

	//使用WithValue包装前面的上下文对象ctxb
	oc := otherContext{ctxb}
	ctxc := context.WithValue(oc,"key","andes,pass from main")

	/*
		ctxc = & cancelCtx {
		Context:oc
		}

		同时通过oc.Context找到ctxb.cancelCtx找到ctxa,在ctxa的children字段维护ctxc作为其子节点
	 */

	go workWithValue(ctxc,"work3")
	//故意"sleep" 10秒，让work2、work3超时退出
	time.Sleep(10 * time.Second)

	//显示调用work1的cancel方法通知其退出
	cancel()
	//等待work1打印退出信息
	time.Sleep(5 * time.Second)
	fmt.Println("main stop")
}

func work(ctx context.Context,name string)  {
	for  {
		select {
		case <- ctx.Done():
			fmt.Printf("%s get msg to cancel \n",name)
			return
		default:
			fmt.Printf("%s is running \n",name)
			time.Sleep(1 * time.Second)
		}
	}
}

//等待前端的退出通知，并试图获取Context传递的数据
func workWithValue(ctx context.Context,name string)  {
	for  {
		select {
		case <- ctx.Done():
			fmt.Printf("%s get msg to cancel \n",name)
			return
		default:
			value := ctx.Value("key").(string)
			fmt.Printf("%s is running value = %s \n",name,value)
			time.Sleep(1 * time.Second)
		}
	}
}