package main

//func main() {
//	for i := 0; i < 3; i++ {
//		i := i //定义一个循环体内局部变量i
//		defer func() {
//			println(i) // 2 1 0
//		}()
//	}
//}

func main() {
	for i := 0; i < 3; i++ {
		defer func(i int) {
			//通过函数传入i
			//defer 语句会马上对调用参数求值
			println(i) // 2 1 0
		}(i)
	}
}
