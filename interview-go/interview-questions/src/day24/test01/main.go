package main

import "fmt"

func main() {
	m := map[int]string{
		0:"zoro",
		1:"one",
	}

	for k,v := range m {
		fmt.Println(k,v)
	}
}
