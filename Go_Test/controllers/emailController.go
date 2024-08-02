package controllers

import (
	"fmt"
	"go_test/config"

	"gopkg.in/gomail.v2"
)

// Send email using gomail
func SendEmail(name string, to string, subject string, body string) {
	// Grab config variables
	cfg := config.LoadConfig()

	m := gomail.NewMessage()

	// Set E-Mail sender with name
	m.SetAddressHeader("From", cfg.SMTPAuthEmail, name)

	// Set E-Mail receivers
	m.SetHeader("To", to)

	// Set E-Mail subject
	m.SetHeader("Subject", subject)

	// Set E-Mail body
	m.SetBody("text/plain", body)

	// Settings for SMTP server
	d := gomail.NewDialer(cfg.SMTPServer, cfg.SMTPPort, cfg.SMTPAuthEmail, cfg.SMTPAuthPassword)

	// This is only needed if SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	//d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Email sent successfully")
}
