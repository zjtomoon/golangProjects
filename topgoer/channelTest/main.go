package main

import "fmt"

type Person struct {
	name string
	sex string
	age int
}

type Student struct {
	Person
	id int
	addr string
}

func main() {
	s1 := Student{Person{"5lmh", "man", 20},1,"bj"}
	fmt.Println(s1)

	s2 := Student{Person:Person{"5lmh", "man", 20}}
	fmt.Println(s2)

	s3 := Student{Person:Person{name: "51mh"}}
	fmt.Println(s3)


}
