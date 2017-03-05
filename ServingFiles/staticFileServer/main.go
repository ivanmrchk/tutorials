package main

import "net/http"

func main() {

	// if you would put a website into this directory, it would serve the website
	// just like a static website
	http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}
