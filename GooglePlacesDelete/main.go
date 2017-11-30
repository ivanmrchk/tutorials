package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type deletePlace struct {
	PlaceId string `json:"place_id"`
}

func check(err error) {
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
}
func main() {
	// api := "AIzaSyBqZQfglwk4iAaO3RT51nEhnuyE3hZXNrY"
	requestUrl := "https://maps.googleapis.com/maps/api/place/delete/json?key=AIzaSyBL_7YvhKHVVdThonbVenBEJ92emxL4TUo"

	obj := deletePlace{
		PlaceId: "qgYvCi0wMDAwMDAzMGYyZjg2ZTU3OjQxM2FhNzEwMDAxOjQyODRhNzJkNTVkZDMwMmE",
	}

	bodyBytes, err := json.Marshal(&obj)
	check(err)
	body := bytes.NewReader(bodyBytes)
	client := &http.Client{}
	req, err := http.NewRequest("POST", requestUrl, body)
	req.Header.Add("Content-Type", "application/json")
	check(err)
	rsp, err := client.Do(req)
	check(err)
	defer rsp.Body.Close()

	body_byte, err := ioutil.ReadAll(rsp.Body)
	check(err)
	fmt.Println(string(body_byte))
}
