package main

import "fmt"

type person struct{
	first string
	last string
	age int
}


	// this is a method, it receives a struct and returns a string.
	// now since this method is receiving the struct person, a variable
	// that has a type of this struct has also access to this method.
func (p person) fullname() string {
	return p.first + p.last
}
func main() {

	// now these variables have access to the fullname() method. because
	// fullname() method has a receiver of person, same as this variable's type
	p1 := person{"James", "Bond", 20}
	p2 := person{"MIss", "Moneypenny", 18 }

	fmt.Println(p1.fullname())
	fmt.Println(p2.fullname())
}
