package main

import (
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

	// This struct has the same field as the person struct,
	// therefore, it will override the person struct field.
type doubleZero struct {
	person
	First         string
	LicenseToKill bool
}

func main() {

	// in both of this variable their first name going to be the first name
	// from doubleZero struct
	p1 := doubleZero{
		person: person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		First:         "Double Zero Seven",
		LicenseToKill: true,
	}

	p2 := doubleZero{
		person: person{
			First: "Miss",
			Last:  "MoneyPenny",
			Age:   19,
		},
		First:         "If looks could kill",
		LicenseToKill: false,
	}

	// fields and methods of the inner-type are promoted to the outer-type
	fmt.Println(p1.First, p1.person.First)
	fmt.Println(p2.First, p2.person.First)
}