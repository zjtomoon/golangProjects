package Context

import (
	"fmt"
	"sync"
	"time"
)

/*
	context库的设计目的就是跟踪goroutinue调用树，并在这些goroutinue调用树中传递通知和元数据
	两个目的:
	(1）退出通知机制——通知可以传递给整个goroutinue调用树的每一个goroutinue。
	(2) 传递数据——数据可以传递给整个goroutinue。
*/
//Context接口
type Context interface {
	//如果Context实现了超时控制，则该方法返回ok,true,deadline为超时时间
	//否则ok为false
	Deadline() (deadline time.Time, ok bool)

	//后端被调的goroutinue应该监听该方法返回的chan,以便及时释放资源
	Done() <-chan struct{}

	//Done返回的chan收到通知的时候，才可以访问Err()获知因什么原因被取消
	Err() error

	//可以访问上游goroutinue传递给下游goroutinue的值
	Value(key interface{}) interface{}
}

//canceler接口
//一个context对象如果实现了canceler接口，则可以被取消
type canceler interface {
	//创建cancel接口实现的goroutinue调用cancel方法通知后续创建的goroutinue退出
	cancel(removeFromParent bool, err error)

	//Done方法返回的chan需要后端goroutinue来监听，并及时退出
	Done() <-chan struct{}
}

//empty Context接口
//emptyCtx实现了Context接口
type emptyCtx int

func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {
	return
}

func (*emptyCtx) Done() <-chan struct{} {
	return nil
}

func (*emptyCtx) Err() error {
	return nil
}

func (*emptyCtx) Value(key interface{}) interface{} {
	return nil
}

func (e *emptyCtx) String() string {
	switch e {
	case background:
		return "context.Background"
	case todo:
		return "context.TODO"
	}
	return "unknown empty Context"
}

var (
	background = new(emptyCtx)
	todo       = new(emptyCtx)
)

func Background() Context {
	return background
}

func TODO() Context {
	return todo
}

func parentCancelCtx(parent Context) (*cancelCtx, bool) {
	for {
		switch c := parent.(type) {
		case *cancelCtx:
			return c, true
		case *timerCtx:
			return &c.cancelCtx, true
		case *valueCtx:
			parent = c.Context
		default:
			return nil, false
		}
	}
}


func removeChild(parent Context, child canceler) {
	p, ok := parentCancelCtx(parent)
	if !ok {
		return
	}
	p.mu.Lock()
	if p.children != nil {
		delete(p.children, child)
	}
	p.mu.Unlock()
}

//cancelCtx
//cancelCtx是一个实现了Context接口的具体类型，同时也实现了canceler接口
//cancelCtx可以被取消，cancelCtx取消时会同时取消所有实现canceler接口的子节点
type cancelCtx struct {
	Context
	done chan struct{} //closed by the first cancel call
	children map[canceler]struct{}
	mu   sync.Mutex
	err  error
}

func (c *cancelCtx) Done() <-chan struct{} {
	return c.done
}

func (c *cancelCtx) Err() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.err
}

func (c *cancelCtx) String() string {
	return fmt.Sprintf("%v.WithCancel", c.Context)
}

func (c *cancelCtx) cancel(removeFromParent bool, err error) {
	if err == nil {
		panic("context:internal error:missing cancel error")
	}
	c.mu.Lock()
	if c.err != nil {
		c.mu.Unlock()
		return
	}
	c.err = err
	//显示通知自己
	close(c.done)

	//循环调用children的cancel函数，由于parent已经取消，所以此时child调用cancel传入的false
	for child := range c.children {
		child.cancel(false, err)
	}
	c, children = nil
	c.mu.Unlock()
	if removeFromParent {
		removeChild(c.Context, c)
	}
}

//timerCtx
/*
	timerCtx是一个实现了Context接口的具体类型，内部封装了cancelCtx类型实例，同时有一个deadline变量，用来实现定时退出通知。

*/

type timerCtx struct {
	cancelCtx
	timer    *time.Time
	deadline time.Time
}

func (c *timerCtx) Deadline() (deadline time.Time, ok bool) {
	return c.deadline, true
}

func (c *timerCtx) String() string {
	return fmt.Sprintf("%v.WithDeadline(%s [%s])", c.cancelCtx.Context, c.deadline, c.deadline.Sub(time.Now()))
}

func (c *timerCtx) cancel(removeFromParent bool, err error) {
	c.cancelCtx.cancel(false,err)
	if removeFromParent {
		removeChild(c.cancelCtx.Context,c)
	}
	c.mu.Lock()
	if c.timer != nil {
		c.timer.Stop()
		c.timer = nil
	}
	c.mu.Unlock()
}


//valueCtx
/*
	value 是一个实现了Context接口具体类型，内部封装了Context接口类型，同时封装了k/v的存储变量。valueCtx用来传递通知信息
*/
type valueCtx struct {
	Context
	key,val interface{}
}

func (c *valueCtx) String() string {
	return fmt.Sprintf("%v.WithiValue(%#v,%#v)",c.Context,c.key,c.val)
}

func (c *valueCtx) Value(key interface{}) interface{} {
	if c.key == key {
		return c.val
	}
	return c.Context.Value(key)
}