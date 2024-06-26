package utilities

import (
	"log"
	"net/smtp"
)

func SendEmail(to string, otp string) error {
	from := "dikocesrt@gmail.com"
	password := "otdl otda orgr ahzp"

	// Konfigurasi SMTP server
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Your OTP Code\n\n" +
		"Your OTP code is " + otp

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return err
	}

	return nil
}