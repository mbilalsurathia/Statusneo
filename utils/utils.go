package utils

import (
	"fmt"
	"log"
	"maker-checker/conf"
	"net/smtp"
)

func SendEmail(emailConfig conf.Email, sendName string, recipient string) error {
	subject := fmt.Sprintf("Subject: Test Email from %v Status Neo", sendName)
	body := "Hello, this is a test email sent from Go!"

	// Combine subject and body
	message := []byte(subject + "\n" + body)

	// Authentication
	auth := smtp.PlainAuth("", emailConfig.From, emailConfig.Password, emailConfig.SmtpHost)

	// Sending the email
	err := smtp.SendMail(emailConfig.SmtpHost+":"+emailConfig.SmtpPort, auth, emailConfig.From, []string{recipient}, message)
	if err != nil {
		log.Fatalf("Failed to send email: %v", err)
	}

	return err
}
