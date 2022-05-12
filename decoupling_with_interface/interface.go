package main

import "fmt"

type Walker interface {
	Walk()
}

// Mandate the structs to implement Interface
var _ Walker = new(Person)
var _ Walker = (*Animal)(nil)

type Person struct{}

func (p Person) Walk() {
	fmt.Println("Person Walking...")
}

type Animal struct{}

func (a Animal) Walk() {
	fmt.Println("Animal Walking...")
}

func main() {
	var walker Walker
	walker = Person{}
	walker.Walk()

	walker = Animal{}
	walker.Walk()
}
