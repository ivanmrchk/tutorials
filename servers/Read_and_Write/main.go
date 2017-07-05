package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
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

		go handle(conn)
	}

}

func handle(conn net.Conn) {

	// conn returns a reader, so we can use that in scanner
	scanner := bufio.NewScanner(conn)

	// scanner.Scan return a bool so unless there is an error in scan or end of
	// document, by default scanner.Scan() scans token by token which are a line.
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		// conn is a writter, so here we are going to write to the connection the
		// line we just read
		fmt.Fprintf(conn, "I headr yousay: %s\n", ln)
	}
	defer conn.Close()

	fmt.Println("Code got here.")
}
