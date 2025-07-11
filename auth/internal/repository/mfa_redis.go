package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type MFARedisRepository struct {
	client *redis.Client
}

func NewMFARedisRepository(client *redis.Client) *MFARedisRepository {
	return &MFARedisRepository{client: client}
}

func (r *MFARedisRepository) SetMFAOtp(ctx context.Context, email string, otp string) error {
	err := r.client.Set(ctx, fmt.Sprintf("mfaotp:%s", email), otp, 120*time.Second)

	if err != nil {
		return fmt.Errorf("failed to set mfa otp in Redis, cause: %s", err)
	}

	return nil
}

func (r *MFARedisRepository) GetMFAOtpByEmail(ctx context.Context, email string) (*string, error) {
	var code string
	err := r.client.Get(ctx, fmt.Sprintf("mfaotp:%s", email)).Scan(&code)

	if err != nil {
		return nil, fmt.Errorf("failed to get mfa otp in Redis, cause: %s", err)
	}

	return &code, nil
}
