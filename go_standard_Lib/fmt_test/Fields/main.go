package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Fields are %q", strings.Fields("  fool bar  baz  "))
}

//Fields are ["fool" "bar" "baz"]
