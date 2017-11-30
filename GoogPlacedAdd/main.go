package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type latlng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type newPlace struct {
	Location    latlng    `json:"location"`
	Accuracy    int       `json:"accuracy"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	Types       [1]string `json:"types"`
}

func check(err error) {
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
}
func main() {
	// api := "AIzaSyBqZQfglwk4iAaO3RT51nEhnuyE3hZXNrY"
	requestUrl := "https://maps.googleapis.com/maps/api/place/add/json?key=AIzaSyBL_7YvhKHVVdThonbVenBEJ92emxL4TUo"

	types := [1]string{"storage"}
	obj := newPlace{
		Location: latlng{
			Lat: 52.1502824,
			Lng: 38.2643063,
		},
		Name:  "01",
		Types: types,
	}

	bodyBytes, err := json.Marshal(&obj)
	check(err)
	body := bytes.NewReader(bodyBytes)
	client := &http.Client{}
	req, err := http.NewRequest("POST", requestUrl, body)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		fmt.Println(err)
	}
	rsp, err := client.Do(req)
	check(err)
	defer rsp.Body.Close()

	body_byte, err := ioutil.ReadAll(rsp.Body)
	check(err)
	fmt.Println(string(body_byte))
}
