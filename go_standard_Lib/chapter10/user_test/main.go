package main

import (
	"fmt"
	"os/user"
)

func main() {
	fmt.Println(user.Current())
	fmt.Println(user.Lookup("xuqingeng"))
	fmt.Println(user.LookupId("0"))
}
