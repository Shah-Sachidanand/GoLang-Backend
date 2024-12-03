package repository

import (
	"context"
	"log"
	"time"

	"learning-golang/app/config"
	"learning-golang/app/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser(user *models.User) (*mongo.InsertOneResult, error)
	GetUsers() ([]*models.User, error)
	GetUserByID(id string) (*models.User, error)
	UpdateUser(id string, user *models.User) (*mongo.UpdateResult, error)
	DeleteUser(id string) (*mongo.DeleteResult, error)
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(client *config.Resource) UserRepository {
	collection := client.DB.Collection("users")
	return &userRepository{
		collection: collection,
	}
}

func (r *userRepository) CreateUser(user *models.User) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return r.collection.InsertOne(ctx, user)
}

func (r *userRepository) GetUsers() ([]*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []*models.User
	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			log.Println(err)
			continue
		}
		users = append(users, &user)
	}

	return users, cursor.Err()
}

func (r *userRepository) GetUserByID(id string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user models.User
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) UpdateUser(id string, user *models.User) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	update := bson.M{
		"$set": user,
	}

	return r.collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
}

func (r *userRepository) DeleteUser(id string) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return r.collection.DeleteOne(ctx, bson.M{"_id": objID})
}
