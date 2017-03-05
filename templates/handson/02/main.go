package main

import (
	"html/template"
	"log"
	"os"
)

type City struct {
	Name       string
	Population int
}

type Country struct {
	Name    string
	Capital string
	Cities  []City
}

type Countries []Country

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	c := Countries{
		Country{
			Name:    "Russia",
			Capital: "Moscow",
			Cities: []City{
				City{
					Name:       "St. Peterburn",
					Population: 4,
				},
				City{
					Name:       "NN",
					Population: 2,
				},
			},
		},
		Country{
			Name:    "Ukraine",
			Capital: "Kiev",
			Cities: []City{
				City{
					Name:       "Odessa",
					Population: 4,
				},
				City{
					Name:       "Lvov",
					Population: 1,
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, c)
	if err != nil {
		log.Fatalln(err)
	}

}
