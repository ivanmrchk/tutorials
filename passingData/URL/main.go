package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/search", foo)
	http.HandleFunc("/", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, v)
}

func bar(w http.ResponseWriter, req *http.Request) {
	///v := req.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
  <form method="GET" action="/search">
   <input type="text" name="q">
   <input type="submit">
  </form>
  <br>`)
}
