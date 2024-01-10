package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//var a int
	//var b string
	//for {
	//	fmt.Println("please input a:")
	//	_,err := fmt.Scanln(&a)
	//	if err != nil {
	//		break
	//	}
	//	fmt.Println("please input b:")
	//	_,err = fmt.Scanln(&b)
	//	if err != nil {
	//		break
	//	}
	//	fmt.Println(a,b)
	//}

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		a := input.Text()
		input.Scan()
		b := input.Text()
		fmt.Println(a, b)
	}

}
