package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type MfaRedisRepository struct {
	client *redis.Client
}

func NewMfaRedisRepository(client *redis.Client) *MfaRedisRepository {
	return &MfaRedisRepository{client: client}
}

func (r *MfaRedisRepository) SetMfaOtp(ctx context.Context, email string, otp string) error {
	err := r.client.Set(ctx, fmt.Sprintf("mfaotp:%s", email), otp, 120*time.Second).Err()

	if err != nil {
		return fmt.Errorf("failed to set mfa otp in Redis, cause: %s", err)
	}

	return nil
}

func (r *MfaRedisRepository) GetMfaOtpByEmail(ctx context.Context, email string) (*string, error) {
	code, err := r.client.Get(ctx, fmt.Sprintf("mfaotp:%s", email)).Result()

	if err != nil {
		return nil, fmt.Errorf("failed to get mfa otp in Redis, cause: %s", err)
	}

	return &code, nil
}
