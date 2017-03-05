package main

import (
	"log"
	"os"
)

func main() {

	// create file
	f, err := os.Create("hello.txt")

	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	// optional you can do this:
	// str := "txt"
	// bs := []byte(str)
	// f.Write(bs)

	// this will write bytes to file
	_, err = f.Write([]byte("text"))
	if err != nil {
		log.Fatalln("Error writing to file:", err.Error())
	}
}
