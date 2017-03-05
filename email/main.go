package main

import (
	"bytes"
	"html/template"
	"log"

	gomail "gopkg.in/gomail.v2"
)

type information struct {
	Name   string
	SendTo string
}

func (i information) sendMail(fromEmail string) {

	t := template.New("template.html")

	var err error
	t, err = t.ParseFiles("template.html")
	if err != nil {
		log.Println(err)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, i); err != nil {
		log.Println(err)
	}

	result := tpl.String()
	m := gomail.NewMessage()
	m.SetHeader("From", fromEmail)
	m.SetHeader("To", "imarchenko@gmail.com")
	m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Email from "+fromEmail)
	m.SetBody("text/html", result)
	m.Attach("template.html")

	d := gomail.NewDialer("smtp.gmail.com", 587, "imarchenko", "iwOZ0?ezOL8,")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
func main() {

	d := information{"jack", "from@from.from"}
	fromE := d.SendTo
	d.sendMail(fromE)
}
