package domain

type SignUpRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Password string `json:"password" validate:"required,min=6,max=100"`
	Email    string `json:"email" validate:"required,email"`
}

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

type SignInResponse struct {
	IsMfaRequired bool             `json:"is_2fa_required"`
	Session       *SessionResponse `json:"session,omitempty"`
}

type SubmitMfaOtpRequest struct {
	Email string `json:"email" validate:"required,email"`
	Otp   string `json:"otp" validate:"required,len=6"`
}

type SessionResponse struct {
	AccessToken        string `json:"access_token"`
	RefreshToken       string `json:"refresh_token"`
	AccessTokenExpiry  int64  `json:"access_token_expiry"`
	RefreshTokenExpiry int64  `json:"refresh_token_expiry"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type RefreshTokenResponse struct {
	AccessToken        string `json:"access_token"`
	AccessTokenExpiry  int64  `json:"access_token_expiry"`
	RefreshToken       string `json:"refresh_token"`
	RefreshTokenExpiry int64  `json:"refresh_token_expiry"`
}
