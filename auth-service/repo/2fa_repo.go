package repo

import (
	"context"
)

type TfaRepo interface {
	Set2FAOtp(ctx context.Context, email string, otp string) error
	Get2FAOtpByEmail(ctx context.Context, email string) (*string, error)
}
