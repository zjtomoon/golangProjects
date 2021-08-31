package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("The num of environ:", len(os.Environ()))
	godebug, ok := os.LookupEnv("GODEBUG")
	if ok {
		fmt.Println("GODEBUG==", godebug)
	} else {
		fmt.Println("GODEBUG not exists")
		os.Setenv("GODEBUG", "gctrace=1")
		fmt.Println("after setenv", os.Getenv("GODEBUG"))
	}
	os.Clearenv()
	fmt.Println("clearenv,the num:", len(os.Environ()))
}
