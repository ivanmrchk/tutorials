package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

	// net.Listen asks for two things, what do you want to listed for which is TCP
	// and what port do you want to listed to that on.
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		io.WriteString(conn, "\nHellow from TPC server \n")
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v", "Well, I hope!")
	}
}
