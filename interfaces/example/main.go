package main

import (
	"fmt"
	"math"
)

type GoQuote struct {
	side float64
}

type GoContact struct {
	radius float64
}
type sendEmail interface {
	area() float64
}

// Anything that has a type of square can use this method
func (q GoQuote) area() float64 {
	return q.side * q.side
}

// Anything that has a type of circle can use this method
func (c GoContact) area() float64 {
	return math.Pi * c.radius * c.radius
}

// info() can take cirlce or a squer since both of them implement area() method.
// And area() method is of interface shape.

//Since info is taking interface, and interface has a method of area.
// anything that has a method of area, can be sent through this function.
func info(e sendEmail) {
	fmt.Println(e)
	fmt.Println(e.area())
}

func main() {
	q := GoQuote{10}
	c := GoContact{5}

	info(q)
	info(c)

}

// Interfaces has a method, and that method is implemented by the struct,
// then any function that receives interface as an arguments, has access
// not only to the method but also to the struct. Meaning, the function with
// an interface can receive a struct as an argument.

/*
1. info( square struct )
2. info() { area() method which is implemented by a struct }
3. Since area is implemented by the struct and info has an argument of interface.
   The info function has an access to struct's method

Depending what is passed into info, it runs different code.


*/
