package domain

import (
	"time"

	"github.com/google/uuid"
)

type AccountResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type UpdateAccount struct {
	Name string `json:"name" validate:"required,min=1,max=100"`
}
