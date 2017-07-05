package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/bluele/gforms"
)

type User struct {
	Name    string  `gforms:"name"`
	Tell    float32 `gforms:"telephone"`
	Email   string  `gforms:"email"`
	Company string  `gforms:"company"`
}

func home(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/index.gohtml"))
	w.Header().Set("Content-Type", "text/html")
	// deine new form = userForm
	quoteFrom := gforms.DefineForm(gforms.NewFields(
		gforms.NewTextField(

			"name",

			gforms.Validators{
				gforms.Required(),
				gforms.MaxLengthValidator(32),
			},
		),
		gforms.NewTextField(
			"Telephone",
			gforms.Validators{},
		),
		gforms.NewTextField(
			"email",
			gforms.Validators{
				gforms.Required("this is needed"),
				gforms.EmailValidator("please enter normal email"),
			},
		),
		gforms.NewTextField(
			"company",
			gforms.Validators{
				gforms.Required("tell us about your company"),
			},
		),
	))

	form := quoteFrom(r)

	if r.Method != http.MethodPost {
		err := tpl.ExecuteTemplate(w, "index.gohtml", form)
		if err != nil {
			log.Fatalln(err)
		}
		return
	}

	// if not validated pass in errors
	if !form.IsValid() {
		err := tpl.ExecuteTemplate(w, "index.gohtml", form)
		if err != nil {
			log.Fatalln(err)
		}
		return
	}

	// maps struct to form
	user := User{}
	form.MapTo(&user)
	fmt.Fprintf(w, "ok: %v", user)
	fmt.Fprintln(os.Stdout, user)
}

func main() {

	http.HandleFunc("/", home)

	http.ListenAndServe(":9000", nil)

}
