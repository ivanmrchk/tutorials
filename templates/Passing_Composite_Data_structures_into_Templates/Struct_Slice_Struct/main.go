package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}
type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	b := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	g := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	m := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	// thi si a slice of structs
	sages := []sage{b, g, m}

	// this is also a slice of structs
	cars := []car{f, c}

	// this will pass in a struct which has two variables (wisdom and transport), those variable are slices, that have structs with in them.
	data := struct {
		Wisdom    []sage
		Transport []car
		//oneMoreType and its slice
	}{
		sages,
		cars,
		// name of slice that has structs in it.
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
