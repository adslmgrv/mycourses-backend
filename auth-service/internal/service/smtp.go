package service

import (
	"fmt"
	"net/smtp"
)

type SmtpEmailService struct {
	Host     string
	Port     uint16
	Username string
	Password string
	From     string
}

func NewSmtpEmailService(host string, port uint16, username, password, from string) *SmtpEmailService {
	return &SmtpEmailService{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		From:     from,
	}
}

func (s *SmtpEmailService) sendEmail(to, subject, body string) error {
	auth := smtp.PlainAuth("", s.Username, s.Password, s.Host)

	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-version: 1.0;\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\r\n" +
		"\r\n" +
		body + "\r\n")

	err := smtp.SendMail(s.Host+":"+string(rune(s.Port)), auth, s.From, []string{to}, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}

func (s *SmtpEmailService) SendSignUpMFAEmail(to, otp string) error {
	subject := "Your MFA Sign-Up Code"
	body := "<p>Thank you for signing up! Your MFA code is: <strong>" + otp + "</strong></p>"
	return s.sendEmail(to, subject, body)
}

func (s *SmtpEmailService) SendSignInMFAEmail(to, otp string) error {
	subject := "Your MFA Sign-In Code"
	body := "<p>Your MFA sign-in code is: <strong>" + otp + "</strong></p>"
	return s.sendEmail(to, subject, body)
}
