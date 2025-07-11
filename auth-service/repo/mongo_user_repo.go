package repo

import (
	"context"

	"github.com/adslmgrv/mycourses-backend/auth-service/error"
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

func (r *MongoUserRepo) FindByUsername(ctx context.Context, username string) (*model.User, error.Error) {
	var user model.User

	if err := r.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, error.Errorf(error.InternalError, "failed to find user by username: %w", err)
	}

	return &user, nil
}

func (r *MongoUserRepo) CreateUser(ctx context.Context, user *model.User) error.Error {
	_, err := r.collection.InsertOne(ctx, user)

	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return error.Errorf(error.EmailTakenError, "email is already taken")
		}
		return error.Errorf(error.InternalError, "failed to create user: %w", err)
	}

	return nil
}
