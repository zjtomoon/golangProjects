package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%q\n", strings.Split("foo,bar,baz", ","))      //Split 会将 s 中的 sep 去掉
	fmt.Printf("%q\n", strings.SplitAfter("foo,bar,baz", ",")) //SplitAfter 会保留 sep
	fmt.Printf("%q\n", strings.SplitN("foo,bar,baz", ",", 2))

	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
}

//["foo" "bar" "baz"]
//["foo," "bar," "baz"]
//["foo" "bar,baz"]

//["a" "b" "c"]
//["" "man " "plan " "canal panama"]
//[" " "x" "y" "z" " "]
//[""]
