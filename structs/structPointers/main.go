package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	p1 := &person{"james", 20}
	fmt.Println(p1)
}