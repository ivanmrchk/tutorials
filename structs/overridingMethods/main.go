package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

	// this struct is embedding person struct.
type doubleZero struct {
	person
	LicenseToKill bool
}
	// this is a method of person struct, and both struct have this struct
func (p person) Greeting() {
	fmt.Println("I'm just a regular person.")
}
	// this method is for doubleZero struct
func (dz doubleZero) Greeting() {
	fmt.Println("Miss Moneypenny, so good to see you.")
}

func main() {

	// when Greeting method is called on this variable it will print I'm just a regular person
	p1 := person{
		Name: "Ian Flemming",
		Age:  44,
	}
	// even thou this variable has a person struct imbedded in it and that struct has a Greeting()
	// method the Greeting method that will be used is of declared type, which is doubleZero.
	p2 := doubleZero{
		person: person{
			Name: "James Bond",
			Age:  23,
		},
		LicenseToKill: true,
	}
	p1.Greeting()
	p2.Greeting()
	p2.person.Greeting()
}