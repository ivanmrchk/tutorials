package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

func main() {

	// this will read the connection 9000
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}
	bs, _ := ioutil.ReadAll(conn)
	fmt.Println(string(bs))
}
