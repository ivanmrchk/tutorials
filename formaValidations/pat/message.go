package main

import (
	"regexp"
	"strings"
)

type GetQuote struct {
	Name        string
	Email       string
	Phone       string
	CompanyName string
	Message     string
	Errors      map[string]string
}

func (q *GetQuote) Validate() bool {

	// create a map for errors
	q.Errors = make(map[string]string)

	// check for errors and populate the struct
	re := regexp.MustCompile(".+@.+\\..+")
	matched := re.Match([]byte(q.Email))

	if strings.TrimSpace(q.Name) == "" {
		q.Errors["Name"] = "Please enter your name"
	}

	if matched == false {
		q.Errors["Email"] = "Please enter a valid email address"
	}

	if strings.TrimSpace(q.Phone) == "" {
		q.Errors["Phone"] = "Please enter your phone number"
	}

	if strings.TrimSpace(q.CompanyName) == "" {
		q.Errors["CompanyName"] = "Please enter company you work for"
	}

	if strings.TrimSpace(q.Message) == "" {
		q.Errors["Message"] = "Please write a message"
	}

	return len(q.Errors) == 0
}
