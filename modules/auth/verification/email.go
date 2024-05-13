package verification

import (
	"fmt"
	"net/smtp"
)

var (
	authUserName   = "AKIAWYDZ23IFJVUH7XYW"
	authPassword   = "BDbf9Zg/czMXycZgrm4frwxYjvFQYUyUbtkBvmXaeA55"
	smtpServerAddr = "email-smtp.ap-south-1.amazonaws.com"
	smtpServerPort = "587"
	senderEmail    = "auth@mail.shivamjha.com"
)

func SendMail(email string, otp string) error {
	fmt.Println("sending emails example")

	msg := []byte("Subject: OTP Email\r\n" +
		"\r\n" +
		fmt.Sprintf("Here is your OTP: %s \r\n", otp))

	auth := smtp.PlainAuth("", authUserName, authPassword, smtpServerAddr)

	return smtp.SendMail(smtpServerAddr+":"+smtpServerPort, auth, senderEmail, []string{email}, msg)
}
