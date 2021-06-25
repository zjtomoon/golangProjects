package main

import (
	"math/rand"
	"fmt")


type Job struct {
	Id      int
	RandNum int
}

type Result struct {
	job *Job
	sum int
}
func main() {
	//需要2个通道
	//1 job通道
	jobChan := make(chan *Job, 128)
	//2 结果通道
	resultChan := make(chan *Result, 128)
	//创建工作池
	createPool(64, jobChan, resultChan)
	//开启打印的协程
	go func(resultChan chan *Result) {
		for result := range resultChan {
			fmt.Printf("job id:%v,randnun:%v,result:%v\n",result.job.Id,result.job.RandNum,result.sum)
		}
	}(resultChan)
	var id int
	//循环创建job,输入到管道
	for {
		id++
		//生成随机数
		r_num := rand.Int()
		job := &Job{
			Id: id,
			RandNum: r_num,
		}
		jobChan <- job
	}
}
//创建工作池
//参数1 ： 开几个协程
func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	//根据开协程个数，去跑运行
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			//执行运算
			//遍历job管道所有数据，进行相加
			for job := range jobChan {
				r_num := job.RandNum
				//随机数每一位相加
				//定义返回值
				var sum int
				for r_num != 0 {
					tmp := r_num % 10
					sum += tmp
					r_num /= 10
				}
				//想要的结果是Result
				r := &Result{
					job: job,
					sum: sum,
				}
				resultChan <- r
			}

		}(jobChan, resultChan)
	}
}

