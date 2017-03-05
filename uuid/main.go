package main

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {

	// request cookie return a pointer to a cookie and an error
	cookie, err := r.Cookie("session")

	// if there is an error, it means there is no cookie
	if err != nil {
		id := uuid.NewV4()
		cookie = &http.Cookie{
			Name: "session",

			// string method converts uuid to string
			Value: id.String(),

			// HTTPS
			// secure: true,

			// this will not allow JavaScript to access it
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)
}
