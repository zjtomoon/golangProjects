package main

import "gopkg.in/macaron.v1"

func main() {
	m := macaron.Classic()

	m.Get("/", func() string{
		return "hello world !"
	})
	m.Run()
}
