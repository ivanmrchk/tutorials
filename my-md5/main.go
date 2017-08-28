package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func md5file(fileName string) []byte {

	// opens a file
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	// defec close (make sure that the file will get closed)
	defer f.Close()

	// create a new hash.
	h := md5.New()
	io.Copy(h, f)
	return h.Sum(nil)
}
func main() {
	// this will walk the filepath which is "."
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		// this path variable is from walk func I can also use info.Name()
		bs := md5file(path)
		fmt.Printf("%x\n", bs)
		return nil
	})

}
