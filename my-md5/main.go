package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	defer f.Close()

	h := sha256.New()

	// copy can be called on fnv.New64, because fnv.New64 embeds a hash interface
	// which embeds a writer interface.
	io.Copy(h, f)

	// h.Sum(nil) - nil there is because you can give it some data
	fmt.Printf("%x", h.Sum(nil))

}
