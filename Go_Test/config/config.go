package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config struct {
	DBURI            string
	DBName           string
	Port             int
	SMTPServer       string
	SMTPPort         int
	SMTPAuthEmail    string
	SMTPAuthPassword string
}

func LoadConfig() Config {
	os.Setenv("DB_URI", "mongodb://localhost:27017/")
	os.Setenv("DB_NAME", "GO_DB")
	os.Setenv("PORT", "8080")
	os.Setenv("SMTP_SERVER", "smtp.gmail.com")
	os.Setenv("SMTP_PORT", "587")
	os.Setenv("AUTH_EMAIL", "")
	os.Setenv("AUTH_PASSWORD", "")

	log.Println("DB_URI:", os.Getenv("DB_URI"))
	log.Println("DB_NAME:", os.Getenv("DB_NAME"))
	log.Println("PORT:", os.Getenv("PORT"))
	log.Println("SMTP_SERVER:", os.Getenv("SMTP_SERVER"))
	log.Println("SMTP_PORT:", os.Getenv("SMTP_PORT"))
	log.Println("AUTH_EMAIL:", os.Getenv("AUTH_EMAIL"))
	log.Println("AUTH_PASSWORD:", os.Getenv("AUTH_PASSWORD"))

	smtpPortInt, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		fmt.Println(err)
	}

	hostPortInt, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		fmt.Println(err)
	}

	return Config{
		DBURI:            os.Getenv("DB_URI"),
		DBName:           os.Getenv("DB_NAME"),
		Port:             hostPortInt,
		SMTPServer:       os.Getenv("SMTP_SERVER"),
		SMTPPort:         smtpPortInt,
		SMTPAuthEmail:    os.Getenv("AUTH_EMAIL"),
		SMTPAuthPassword: os.Getenv("AUTH_PASSWORD"),
	}
}
