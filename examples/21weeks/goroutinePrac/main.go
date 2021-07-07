package main

/*
   使用goroutine和channel实现一个计算int64随机数各位数和的程序
   1、开启一个goroutine循环生成int64类型的随机数，发送到jobChan
   2、开启24个goroutine从jobChan中取出
   3、主goroutine从resultChan取出结果并打印到终端输出
 */
func main() {
	
}
