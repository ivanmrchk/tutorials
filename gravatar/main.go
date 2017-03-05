package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"strings"
)

// This function takes an email strips it off space a
// and converts is to lower space, to get the email ready
// hasshing.
func getGravatarHash(email string) string {
	email = strings.TrimSpace(email)
	email = strings.ToLower(email)

	// som
	// stuff
	h := md5.New()
	io.WriteString(h, email)

	finalBytes := h.Sum(nil)

	finalString := "http://www.gravatar.com/avatar/" + hex.EncodeToString(finalBytes) + "?d=identicon"

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
