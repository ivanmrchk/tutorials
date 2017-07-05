package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {

	// os.Args takes an argument from terminal. in this case it will
	// take the second argument, first being the program name.
	//
	// $ go install
	// $ osArgs ../reader/hello.txt
	// ..arg[0] ........2ns arg[1]
	f, err := os.Open(os.Args[1])
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
