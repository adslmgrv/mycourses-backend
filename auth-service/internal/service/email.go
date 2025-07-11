package service

type EmailService interface {
	SendSignUp2FAEmail(to string, otp string) error
	SendSignIn2FAEmail(to string, otp string) error
}
