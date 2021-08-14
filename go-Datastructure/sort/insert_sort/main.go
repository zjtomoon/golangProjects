package main

import (
	"fmt"
	"math/rand"
	"time"
)

func InsertSort(arr *[80000]int) {
	//完成第一次，给第二个元素找到合适的位置并插入

	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]
		insertIndex := i - 1 //下标
		//从大到小
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex] //数据后移
			insertIndex--
		}
		//插入
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}
	}

	/*

		//完成第2次，给第3个元素找到合适的位置并插入
		insertVal = arr[2]
		insertIndex = 2 - 1 // 下标

		//从大到小
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex + 1] = arr[insertIndex] // 数据后移
			insertIndex--
		}
		//插入
		if insertIndex + 1 != 2 {
			arr[insertIndex + 1] = insertVal
		}
		fmt.Println("第2次插入后", *arr)

		//完成第3次，给第4个元素找到合适的位置并插入
		insertVal = arr[3]
		insertIndex = 3 - 1 // 下标

		//从大到小
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex + 1] = arr[insertIndex] // 数据后移
			insertIndex--
		}
		//插入
		if insertIndex + 1 != 3 {
			arr[insertIndex + 1] = insertVal
		}
		fmt.Println("第3次插入后", *arr)

		//完成第4次，给第5个元素找到合适的位置并插入
		insertVal = arr[4]
		insertIndex = 4 - 1 // 下标

		//从大到小
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex + 1] = arr[insertIndex] // 数据后移
			insertIndex--
		}
		//插入
		if insertIndex + 1 != 4 {
			arr[insertIndex + 1] = insertVal
		}
		fmt.Println("第4次插入后", *arr)*/
}

func main() {
	var arr [80000]int
	for i := 0; i < 80000; i++ {
		arr[i] = rand.Intn(900000)
	}

	start := time.Now()
	InsertSort(&arr)
	end := time.Now()
	fmt.Println(arr)
	fmt.Println("插入排序耗时 = ",end.Sub(start))
	fmt.Println("main函数")
}