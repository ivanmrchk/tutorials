package main

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	UserName, First, Last string
}

var tpl *template.Template

// this map will hold all the users.
var dbUsers = map[string]user{}

// this map will hold user id and session id
var dbSessions = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	// get cookie
	c, err := r.Cookie("session")
	if err != nil {
		// if there is no cookie session create one
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name: "session",
			// convert uuid to string
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	} // error handling

	var u user

	// basically what this does, is check if cookie session has a key value and if
	// that value is not blank.
	if un, ok := dbSessions[c.Value]; ok {
		// If it is not blank (, ok idiom) it will assign the c.Value to dbUsers as a
		// key un and a value the c.Value
		u = dbUsers[un]
	}

	// Process Form
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")
		u = user{un, f, l}

		// seession id or c.Value is equal to user name in the dbSessions map
		dbSessions[c.Value] = un

		// un key is equal to user struct in dbUsers map
		dbUsers[un] = u
	}

	tpl.ExecuteTemplate(w, "index.gohtml", u)

}

func bar(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	// if there is no cookie seesion go back to log in
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// grab the session id and check if it is there.
	un, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// if it got here, then it means that  dbSessions[c.Value] is not empty

	// create a struct of user information
	// this un is a uuid from cookie
	u := dbUsers[un]
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}
