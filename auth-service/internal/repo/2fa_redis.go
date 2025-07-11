package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis2FARepo struct {
	client *redis.Client
}

func NewRedis2FARepo(client *redis.Client) *Redis2FARepo {
	return &Redis2FARepo{client: client}
}

func (r *Redis2FARepo) Set2FAOtp(ctx context.Context, email string, otp string) error {
	err := r.client.Set(ctx, fmt.Sprintf("2fa_otp:%s", email), otp, 120*time.Second)

	if err != nil {
		return fmt.Errorf("Failed to set 2fa otp in Redis, cause: %s", err)
	}

	return nil
}

func (r *Redis2FARepo) Get2FAOtpByEmail(ctx context.Context, email string) (*string, error) {
	var code string
	err := r.client.Get(ctx, fmt.Sprintf("2fa_otp:%s", email)).Scan(&code)

	if err != nil {
		return nil, fmt.Errorf("Failed to get 2fa otp in Redis, cause: %s", err)
	}

	return &code, nil
}
