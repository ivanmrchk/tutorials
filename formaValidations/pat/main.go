package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/bmizerany/pat"
)

func main() {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(index))
	mux.Post("/", http.HandlerFunc(send))
	mux.Get("/confirmation", http.HandlerFunc(confirmation))

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}

func index(w http.ResponseWriter, r *http.Request) {
	render(w, "templates/index.html", nil)
}

func send(w http.ResponseWriter, r *http.Request) {

	q := &GetQuote{
		Name:        r.FormValue("name"),
		Email:       r.FormValue("email"),
		Phone:       r.FormValue("phone"),
		CompanyName: r.FormValue("companyname"),
		Message:     r.FormValue("message"),
	}

	// if form is not validated pass in error messages.
	if q.Validate() == false {
		render(w, "templates/index.html", q)
		return
	}

	// Send message in an email
	// Redirect to confirmation page
}

func confirmation(w http.ResponseWriter, r *http.Request) {
	render(w, "templates/confirmation.html", nil)
}

func render(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
