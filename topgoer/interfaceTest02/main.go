package main

import "fmt"

type  Sayer interface {
	say()
}
type Mover interface {
	move()
}

type animal interface {
	Sayer
	Mover
}

type cat struct {
	name string
}

func (c cat) say(){
	fmt.Printf("%s:喵喵喵\n",c.name)
}
func (c cat) move()  {
	fmt.Printf("%s is moving\n",c.name)
}

func main()  {
	var x animal
	x = cat{name: "大橘"}
	x.say()
	x.move()
}
