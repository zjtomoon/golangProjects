package main

import "log"

func test(args ...string) { // 可以接受任意个string参数
	for _, v := range args {
		log.Println(v)
	}
}

func main() {
	var arr1 = []string{
		"ele1",
		"ele2",
		"ele3",
	}
	var arr2 = []string{
		"ele4",
		"ele5",
	}
	arr1 = append(arr1, arr2...) // arr2的元素被打散一个个append进arr1
	log.Println(arr1)

	var arr = []string{
		"ele1",
		"ele2",
		"ele3",
	}
	test(arr...) // 切片被打散传入
}

// 总结：
// ... 作为参数表示可以传入多个同类型参数
// ... 还可以用于切片展开，将切片打散后传入
