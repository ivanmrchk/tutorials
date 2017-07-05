package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog")
}

type hotcat int

func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat")
}

func main() {
	var h hotdog
	var c hotcat

	mux := http.NewServeMux()
	mux.Handle("/", h)
	mux.Handle("/", c)

	http.ListenAndServe(":8080", mux)
}
