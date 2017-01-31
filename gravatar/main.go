package main

import (
	"fmt"
	"encoding/hex"
	"crypto/md5"
	"io"
	"strings"
)

func getGravatarHash(email string) string {
	email = strings.TrimSpace(email)
	email = strings.ToLower(email)

	h := md5.New()
	io.WriteString(h, email)


	finalBytes := h.Sum(nil)

	finalString := hex.EncodeToString(finalBytes)


	return finalString
}

func main() {
	gravatarHash := getGravatarHash("aleksmarch.dev@gmail.com")

	fmt.Println(`<!DOCTYPE html>

    <html>
        <head>
        </head>
        <body>
            <img src="http://www.gravatar.com/avatar/` + gravatarHash + `?d=identicon">
        </body>
    </html>

    `)
}