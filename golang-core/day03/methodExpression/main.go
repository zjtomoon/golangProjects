package main

import (
	"fmt"
)

type T struct {
	a int
}

func (t *T) Set(i int) {
	t.a = i
}

func (t *T) Print() {
	fmt.Printf("%p,%v,%d\n", t, t, t.a)
}

func main() {

}
