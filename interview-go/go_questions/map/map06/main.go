package main

import "fmt"

func main() {
	ageMap := make(map[string]int)
	ageMap["tom"] = 18

	//不带comma的用法

	age1 := ageMap["tom"]
	fmt.Println(age1)

	//带comma的用法
	age2,ok := ageMap["tom"]
	fmt.Println(age2,ok)
}
