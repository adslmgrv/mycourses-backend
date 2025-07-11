package service

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/adslmgrv/mycourses-backend/auth-service/dto"
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

func newSixDigitOtp() string {
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}
