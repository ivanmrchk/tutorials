package main

import (
	"log"

	"github.com/hbagdi/go-unsplash/unsplash"
	"golang.org/x/oauth2"
)

func main() {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "b0ed30b19db70ca6377b3905882beaa6505d77418e5fe279194726be54d18d9f"},
	)

	client := oauth2.NewClient(oauth2.NoContext, ts)
	//use the http.Client to instantiate unsplash
	unsplash := unsplash.New(client)

	url, resp, err := unsplash.Photos.DownloadLink("-HPhkZcJQNk")

	log.Println(url, err, resp)

}
