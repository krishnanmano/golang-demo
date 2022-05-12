package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{
		Name: "Terry",
		Age:  29,
	}

	fmt.Println("Before: ", p)
	p.SetName("Tom")
	p.SetAge(100)
	fmt.Println("After: ", p)
}

func (p Person) String() string {
	return fmt.Sprintf("{'name': '%s', 'age': '%d'}", p.Name, p.Age)
}

func (p Person) GetAge() int {
	return p.Age
}

func (p *Person) SetAge(age int) {
	p.Age = age
}

func (p Person) SetName(name string) {
	p.Name = name
}
