package main

import (
	"encoding/binary"
	"fmt"
	"hash/fnv"
	"io"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	defer f.Close()

	h := fnv.New64()

	// copy can be called on fnv.New64, because fnv.New64 embeds a hash interface
	// which embeds a writer interface.
	io.Copy(h, f)

	// h.Sum(nil) - nil there is because you can give it some data
	fmt.Println("Thue sum is", h.Sum64())

	// convert slice of bytes to int
	data := binary.BigEndian.Uint64(h.Sum(nil))
	fmt.Println("Thue sum is", data)
}
