package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template // this variable should be in a package scope, so both functions, init and main could have access to it.

func init() {
	tpl = template.Must(template.ParseGlob("tplDir/*.txt")) // parse glog will parse the entire directory, also must will chekc for error, so no need to check errors. What other thing must does, is  perhibits parsing same file twice.
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "test.txt", nil) // no need to parse any of these templates, since ParseGlob parsed the entire foldf.

	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "text2.txt", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
