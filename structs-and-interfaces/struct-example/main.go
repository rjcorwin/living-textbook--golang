package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	rj := Person{Name: "RJ", Age: 30}
	fmt.Println(rj)
	fmt.Printf("Name: %s\n", rj.Name)
	fmt.Printf("Age: %d\n", rj.Age)
}
