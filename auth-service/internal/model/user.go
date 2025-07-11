package model

import "time"

type User struct {
	ID              string    `bson:"_id,omitempty"`
	Name            string    `bson:"name"`
	Email           string    `bson:"email"`
	IsEmailVerified bool      `bson:"isEmailVerified"`
	PasswordHash    []byte    `bson:"passwordHash"`
	CreatedAt       time.Time `bson:"createdAt"`
}
