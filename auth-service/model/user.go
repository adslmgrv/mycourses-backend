package model

type User struct {
	ID           string `bson:"_id,omitempty"`
	Name         string `bson:"name"`
	Email        string `bson:"email"`
	Username     string `bson:"username"`
	PasswordHash string `bson:"passwordHash"`
	CreatedAt    int64  `bson:"createdAt"`
}
