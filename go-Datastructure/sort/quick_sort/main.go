package main

import (
	"fmt"
	"math/rand"
	"time"
)

//快速排序
//说明
//1. left表示 数组左边的下标
//2. right 表示数组右边的下标
//3. array表示要排序的数组

func QuickSort(left int,right int,array *[80000]int)  {
	l := left
	r := right
	// privot 是中轴，支点
	privot := array[(left + right) /2]
	temp := 0

	//for循环的目标是将比privot小的数放左边
	//比privot大的数放到右边
	for ; l < r ; {
		//从 privot的左边找到大于等于privot的值
		for ; array[l] < privot ;{
			l++
		}
		//从privot的右边找到小于等于privot的值
		for ; array[r] > privot ; {
			r--
		}
		// l >=r 表明本次分解任务完成,break
		if l >= r {
			break
		}
		//交换
		temp = array[l]
		array[l] = array[r]
		array[r] = temp

		//优化
		if array[l] == privot {
			r--
		}
		if array[r] == privot {
			l++
		}
	}
	//如果l ==r 再移动下
	if l ==r {
		l++
		r--
	}
	//向左递归
	if left < r {
		QuickSort(left,r,array)
	}
	//向右递归
	if right > l {
		QuickSort(l,right,array)
	}
}

func main()  {
	// arr := [9]int {-9,78,0,23,-567,70, 123, 90, -23}
	// fmt.Println("初始", arr)
	var arr [80000]int
	for i := 0 ; i < 80000 ; i++ {
		arr[i] = rand.Intn(900000)
	}
	start := time.Now()
	QuickSort(0,len(arr)-1,&arr)
	end := time.Now()
	fmt.Println("main函数")
	fmt.Println(arr)
	fmt.Println("快速排序耗时 = ",end.Sub(start))
}