package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func main() {
	http.Handle("/dog/", http.HandlerFunc(d))

	http.ListenAndServe(":8080", nil)
}
