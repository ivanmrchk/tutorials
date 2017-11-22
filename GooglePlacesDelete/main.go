package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type deletePlace struct {
	PlaceId string `json:"place_id"`
}

func main() {
	// api := "AIzaSyBqZQfglwk4iAaO3RT51nEhnuyE3hZXNrY"
	requestUrl := "https://maps.googleapis.com/maps/api/place/delete/json?key=AIzaSyBL_7YvhKHVVdThonbVenBEJ92emxL4TUo"

	obj := deletePlace{
		PlaceId: "qgYvCi0wMDAwMDAzMGYyZjg2ZTU3OjQxM2FhNzEwMDAxOjQyODRhNzJkNTVkZDMwMmE",
	}

	bodyBytes, err := json.Marshal(&obj)
	if err != nil {
		fmt.Println(err)
	}
	body := bytes.NewReader(bodyBytes)
	client := &http.Client{}
	req, err := http.NewRequest("POST", requestUrl, body)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		fmt.Println(err)
	}
	rsp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()

	body_byte, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body_byte))
}
