package service

type EmailService interface {
	SendSignUpMfaEmail(to string, otp string) error
	SendSignInMfaEmail(to string, otp string) error
}
