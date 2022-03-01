package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	student := make(map[string]interface{})
	student["name"] = "json"
	student["age"] = 18
	student["sex"] = "man"
	b, err := json.Marshal(student)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

	b, err = json.MarshalIndent(student, "", "  ")
	if err != nil {
		fmt.Println("json err", err)
	}
	fmt.Println(string(b))
}
