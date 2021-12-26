package main

func main() {
	for i := 0 ; i < 3; i++ {
		defer func() {
			println(i) // 3 3 3
		}()
	}
}
