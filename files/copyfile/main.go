package main

import (
	"io"
	"log"
	"os"
)

func main() {
	srcFileName := os.Args[1]
	dstFileName := os.Args[2]

	// open file
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		log.Fatalln("Error openning source file: ", err)
	}
	defer srcFile.Close()

	// creating file
	dstFile, err := os.Create(dstFileName)
	if err != nil {
		log.Fatalln("Error creatign a destination file: ", err)
	}

	defer dstFile.Close()

	// or you can also do this
	// _, err = io.Copy(dstFile, string.Newreader("hello world"))

	_, err = io.Copy(dstFile, srcFile)

	if err != nil {
		log.Fatalln("Error writing to destination file: ", err)
	}
}
