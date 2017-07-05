package main

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func getUser(w http.ResponseWriter, r *http.Request) user {
	// get cookie
	c, err := r.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}
	http.SetCookie(w, c)

	// if user exists already, get user
	var u user
	if un, ok := dbSessionsn[c.Value]; ok {
		// if there is information or value in dbSessionsn[c.Value] then assign that
		// value to u and return it
		u = dbUsers[un]
	}
	return u
}

func alreadyLoggedIn(r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}

	// user name is equat to the value of uuid in dbSessionsn
	un := dbSessionsn[c.Value]

	// check to see if there is a user struct under that un
	_, ok := dbUsers[un]
	return ok
}
