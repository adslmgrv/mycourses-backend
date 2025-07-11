package service

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/adslmgrv/mycourses-backend/auth/internal/domain"
	apperr "github.com/adslmgrv/mycourses-backend/auth/internal/error"
	"github.com/adslmgrv/mycourses-backend/auth/internal/model"
	"github.com/adslmgrv/mycourses-backend/auth/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepository repository.UserRepository
	mfaRepository  repository.MfaRepository
	emailService   EmailService
}

func NewAuthService(userRepository repository.UserRepository, mfaRepository repository.MfaRepository, emailService EmailService) AuthService {
	return AuthService{
		userRepository: userRepository,
		mfaRepository:  mfaRepository,
		emailService:   emailService,
	}
}

func (s AuthService) SignUp(ctx context.Context, request domain.SignUpRequest) error {
	user, err := s.userRepository.FindByEmail(ctx, request.Email)
	if err != nil {
		return err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password, cause: %s", err)
	}

	if user != nil && !user.IsEmailVerified {
		err := s.userRepository.UpdatePasswordHashByEmail(ctx, request.Email, passwordHash)

		if err != nil {
			return err
		}

	} else {
		err := s.userRepository.CreateUser(ctx, model.User{
			Name:            request.Name,
			Email:           request.Email,
			IsEmailVerified: false,
			PasswordHash:    passwordHash,
			CreatedAt:       time.Now().UTC(),
		})

		if err != nil {
			return err
		}
	}

	otp := newSixDigitOtp()
	err = s.mfaRepository.SetMFAOtp(ctx, request.Email, otp)
	if err != nil {
		return err
	}

	return s.emailService.SendSignUpMFAEmail(request.Email, otp)
}

func (s AuthService) SubmitMFAOtp(ctx context.Context, request domain.SubmitMFAOtpRequest) (*domain.SessionResponse, error) {
	otp, err := s.mfaRepository.GetMFAOtpByEmail(ctx, request.Email)
	if err != nil {
		return nil, err
	}

	if otp != &request.Otp {
		return nil, apperr.Errorf(apperr.MfaFailedError, "Invalid 2fa otp")
	}

	return nil, nil
}

func newSixDigitOtp() string {
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}
