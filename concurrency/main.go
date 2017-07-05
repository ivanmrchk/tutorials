package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {

		// putting on to a channel
		c <- "ping"
	}
}

func printer(c chan string) {
	for {

		// taking off the channel
		msg := <-c
		fmt.Println(msg)

		time.Sleep(time.Second * 1)
	}
}

func main() {

	// create a channel
	var c chan string = make(chan string)

	go pinger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
