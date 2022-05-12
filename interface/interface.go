package main

import "fmt"

type Walker interface {
	Walk()
}

var _ Walker = new(Person)

type Person struct{}

func (p Person) Walk() {
	fmt.Println("Person Walking...")
}

func main() {
	var walker Walker = Person{}
	walker.Walk()
}
