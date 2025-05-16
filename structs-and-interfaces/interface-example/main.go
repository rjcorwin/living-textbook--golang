package main

import "fmt"

type Greeter interface {
	Greet() string
}

type Person struct {
	Name string
}

func (p Person) Greet() string {
	return "Hello, my name is " + p.Name
}

func sayHello(g Greeter) {
	fmt.Println(g.Greet())
}

func main() {
	rj := Person{Name: "RJ"}
	sayHello(rj)
}
