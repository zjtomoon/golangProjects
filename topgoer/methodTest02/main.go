package main

import "fmt"

type User struct {
	Name  string
	Email string
}

func (u *User) Notify(){
	fmt.Printf("%v : %v \n",u.Name,u.Email)
}

func main() {
	u1 := User{"golang","golang@golang.com"}
	u1.Notify()
	u2 := User{"go","go@go.com"}
	u3 := &u2
	u3.Notify()
}