package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%d\n", strings.IndexFunc("studygolang", func(c rune) bool {
		if c > 'u' {
			return true
		}
		return false
	}))
}

//4
