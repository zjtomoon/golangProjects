package main

import "fmt"

func main() {
	ageMp := make(map[string]int)
	ageMp["tom"] = 18

	for name,age := range ageMp {
		fmt.Println(name,age)
	}
}
