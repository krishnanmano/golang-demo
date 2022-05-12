package main

import (
	"fmt"
)

type Address struct {
	state   string
	country string
}

func (a Address) GetCountry() string {
	return a.country
}

type Person struct {
	name string
	age  int
	Address
}

func (p Person) String() string {
	return fmt.Sprintf("{'name': '%s', 'age': '%d', 'state': '%s', 'country': '%s'}", p.name, p.age, p.state, p.country)
}

func (p Person) GetCountry() string {
	return "india"
}

func main() {
	p := Person{
		name: "Terry",
		age:  29,
		Address: Address{
			state:   "California",
			country: "US",
		},
	}

	fmt.Println(p)
	fmt.Println(p.GetCountry())
	fmt.Println(p.Address.GetCountry())
}
