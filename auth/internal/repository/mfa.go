package repository

import (
	"context"
)

type MfaRepository interface {
	SetMFAOtp(ctx context.Context, email string, otp string) error
	GetMFAOtpByEmail(ctx context.Context, email string) (*string, error)
}
