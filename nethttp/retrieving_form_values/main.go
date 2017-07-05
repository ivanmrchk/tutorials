package main

import (
	"html/template"
	"log"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	// on request it will parse the form. This is needed to get data from form
	req.ParseForm()

	log.Print(req.Form)

	// this will pass in data from req.ParseForm() into a template
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8090", d)
}
