package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string
	Hobby string
}

func main() {
	p := Person{
		"json",
		"games",
	}

	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println("json err", err)
	}
	fmt.Println(string(b))

	// 格式输出
	b, err = json.MarshalIndent(p, "", "  ")
	if err != nil {
		fmt.Println("json err", err)
	}
	fmt.Println(string(b))
}
