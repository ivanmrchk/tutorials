package main

import (
	"io"
	"net"
)

func main() {

	// ln Listener returns an error
	// takes in a protocol and a prt number.
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	// TCP Listener must be defer closed
	defer ln.Close()

	// Allways accept connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		// write to connection
		io.WriteString(conn, "Hello World/n")

		// defer is not used because otherwise it would close the connection at the
		// end of the main function.
		conn.Close()
	}
}

// connection is like a file, you can read and write to it.
