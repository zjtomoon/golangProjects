package ticker

import (
	"time"
	"fmt"
)

func worker()  {
	ticker := time.Tick( 100 * time.Microsecond)
	for {
		select {
		case <- ticker:
			//执行定时任务
			fmt.Println("执行 1s 定时任务")
		}
	}
}

