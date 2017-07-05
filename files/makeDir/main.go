package main

import (
	"log"
	"os"
)

func main() {

	// create dir
	err := os.Mkdir("/abc", 0x777)
	if err != nil {
		log.Fatalln(err)
	}

	// create file
	f, err := os.Create("/abc/hello.txt")

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
