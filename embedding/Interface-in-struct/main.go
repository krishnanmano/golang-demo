package main

import "fmt"

type walker interface {
	walk()
}

type Person struct {
	walker
}

func (p Person) walk() {
	fmt.Println("Person Walking...")
}

func main() {
	p := Person{}
	p.walk()
}
