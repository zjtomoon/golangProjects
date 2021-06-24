package main

import "sync"

func main() {
	wg := sync.WaitGroup{}

	si := []int{1,2,3,4,5,6,7,8,9,10}

	/*for i := range si{
		wg.Add(1)
		go func() {
			println(i)
			wg.Done()
		}()
	}
	wg.Wait()*/
//上面的程序没有遍历切片si,而是全部打印9
	for i := range si {
		wg.Add(1)

		//这里有一个实参到形参的值拷贝
		go func(a int) {
			println(a)
			wg.Done()
		}(i)
	}
	wg.Wait()
//正确的写法是使用函数参数做一次数据复制，而不是闭包。
}
