package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("hello.txt")
	if err != nil {
		log.Fatalln("my program broken")
	}

	defer f.Close()

	rdr := strings.NewReader("some string")
	fmt.Printf("%T", rdr)
	fmt.Println("")

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln("My programm broke")
	}

	str := string(bs)
	fmt.Println(str)

}
