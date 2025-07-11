package service

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/adslmgrv/mycourses-backend/auth-service/dto"
	appe "github.com/adslmgrv/mycourses-backend/auth-service/error"
	"github.com/adslmgrv/mycourses-backend/auth-service/model"
	"github.com/adslmgrv/mycourses-backend/auth-service/repo"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo     repo.UserRepo
	tfaRepo      repo.TfaRepo
	emailService EmailService
}

func NewAuthService(userRepo repo.UserRepo, tfaRepo repo.TfaRepo, emailService EmailService) AuthService {
	return AuthService{
		userRepo:     userRepo,
		tfaRepo:      tfaRepo,
		emailService: emailService,
	}
}

func (s AuthService) SignUp(ctx context.Context, request dto.SignUpRequest) error {
	user, err := s.userRepo.FindByEmail(ctx, request.Email)
	if err != nil {
		return err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password, cause: %s", err)
	}

	if user != nil && !user.IsEmailVerified {
		err := s.userRepo.UpdatePasswordHashByEmail(ctx, request.Email, passwordHash)

		if err != nil {
			return err
		}

	} else {
		err := s.userRepo.CreateUser(ctx, model.User{
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
	err = s.tfaRepo.Set2FAOtp(ctx, request.Email, otp)
	if err != nil {
		return err
	}

	return s.emailService.SendSignUp2FAEmail(request.Email, otp)
}

func (s AuthService) Submit2FAOtp(ctx context.Context, request dto.Submit2FAOtpRequest) (*dto.SessionResponse, error) {
	otp, err := s.tfaRepo.Get2FAOtpByEmail(ctx, request.Email)
	if err != nil {
		return nil, err
	}

	if otp != &request.Otp {
		return nil, appe.Errorf(appe.TfaFailedError, "Invalid 2fa otp")
	}

	return nil, nil
}

func newSixDigitOtp() string {
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}
