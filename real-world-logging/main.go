package main

import (
	"log"
	"net/smtp"
)

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}
func main() {
	// Connect to the remote SMTP server.
	client, err := smtp.Dial("smtp.smail.com:25")
	if err != nil {
		log.Fatalln(err)
	}
	client.Data()
}
