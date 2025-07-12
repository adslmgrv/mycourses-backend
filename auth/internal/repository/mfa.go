package repository

import (
	"context"
)

type MfaRepository interface {
	SetMfaOtp(ctx context.Context, email string, otp string) error
	GetMfaOtpByEmail(ctx context.Context, email string) (*string, error)
}
