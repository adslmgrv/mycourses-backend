package repo

import (
	"context"
	"fmt"

	appe "github.com/adslmgrv/mycourses-backend/auth-service/error"
	"github.com/adslmgrv/mycourses-backend/auth-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepo struct {
	collection *mongo.Collection
}

func NewMongoUserRepo(db *mongo.Database) *MongoUserRepo {
	return &MongoUserRepo{
		collection: db.Collection("users"),
	}
}

func (r *MongoUserRepo) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User

	if err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find user by username: %w", err)
	}

	return &user, nil
}

func (r *MongoUserRepo) UpdatePasswordHashByEmail(ctx context.Context, email string, passwordHash []byte) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"email": email}, bson.M{"$set": bson.M{"passwordHash": passwordHash}})

	if err != nil {
		return fmt.Errorf("failed to update password: %w", err)
	}

	return nil
}

func (r *MongoUserRepo) CreateUser(ctx context.Context, user *model.User) error {
	_, err := r.collection.InsertOne(ctx, user)

	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return appe.Errorf(appe.EmailTakenError, "email is already taken")
		}
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}
