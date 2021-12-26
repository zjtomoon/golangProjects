package main

//func main() {
//	for i:= 0;i<3 ;i++  {
//		i := i
//		defer func() {
//			println(i) // 2 1 0
//		}()
//	}
//}

func main() {
	for i := 0; i < 3; i++ {
		defer func(i int) {
			println(i) // 2 1 0
		}(i)
	}
}
