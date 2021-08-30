package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	Name string
	Age  int
	Sex  int
}

func main() {
	p := &Person{"polaris", 28, 0}
	fmt.Printf("%#v", p)
}

func (this *Person) Gostring() string {
	return "&Person{Name is " + this.Name + ", Age is " + strconv.Itoa(this.Age) + ", Sex is " + strconv.Itoa(this.Sex) + "}"
}
