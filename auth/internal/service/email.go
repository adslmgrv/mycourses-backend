package service

type EmailService interface {
	SendSignUpMFAEmail(to string, otp string) error
	SendSignInMFAEmail(to string, otp string) error
}
