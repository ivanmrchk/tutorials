package main

func main() {

	sender := NewSender("imarchenko@gmail.com", "iwOZ0?ezOL8,")

	//The receiver needs to be in slice as the receive supports multiple receiver
	Receiver := []string{"imarchenko@gmail.com", "xyz@gmail.com", "larrypage@googlemail.com"}

	Subject := "Testing email from golang"
	bodyMessage := "Sending email using Golang. Yeah"

	sender.SendMail(Receiver, Subject, bodyMessage)
}
