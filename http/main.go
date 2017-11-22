package main

import (
	"net/http"

	"github.com/kr/pretty"
)

type MyHandler int

func (h MyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	// fmt.Println(req.Header)

	// settin header
	// res.Header().Set("content-type", "plain/text")

	// write to handler using ResponseWriter
	writeTo := req.URL.Query()
	// io.WriteString(res, writeTo)

	pretty.Println(writeTo)
}

func main() {
	var h MyHandler
	http.ListenAndServe(":8000", h)
}
