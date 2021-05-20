/*package main

import "fmt"

type People interface {
Speak(string) string
}

type Student struct{}

func (stu *Stduent) Speak(think string) (talk string) {
if think == "sb" {
talk = "你是个大帅比"
} else {
talk = "您好"
}
return
}

func main() {
var peo People = Student{}
think := "bitch"
fmt.Println(peo.Speak(think))
}*/
/*package main

import "fmt"

type Cat struct{}

func (c Cat) Say() string { return "喵喵喵" }

type Dog struct{}

func (d Dog) Say() string { return "汪汪汪" }

func main() {
c := Cat{}
fmt.Println("猫:", c.Say())
d := Dog{}
fmt.Println("狗:", d.Say())
}*/

package main

import "fmt"

type Mover interface {
	move()
}

type dog struct {
	name string
}

type car struct {
	brand string
}

func (d dog) move()  {
	fmt.Printf("%s会跑\n",d.name)
}

func (c car) move()  {
	fmt.Printf("%s速度70公里每小时\n",c.brand)
}

func main() {
	var x Mover
	var a = dog{name:"旺财"}
	var b = car{brand: "桑塔纳"}
	x = a
	x.move()
	x = b
	x.move()
}