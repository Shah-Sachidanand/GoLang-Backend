package services

import (
	"learning-golang/app/internal/models"
	"learning-golang/app/internal/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	CreateUser(user *models.User) (*mongo.InsertOneResult, error)
	GetUsers() ([]*models.User, error)
	GetUserByID(id string) (*models.User, error)
	UpdateUser(id string, user *models.User) (*mongo.UpdateResult, error)
	DeleteUser(id string) (*mongo.DeleteResult, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) CreateUser(user *models.User) (*mongo.InsertOneResult, error) {
	return s.repo.CreateUser(user)
}

func (s *userService) GetUsers() ([]*models.User, error) {
	return s.repo.GetUsers()
}

func (s *userService) GetUserByID(id string) (*models.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *userService) UpdateUser(id string, user *models.User) (*mongo.UpdateResult, error) {
	return s.repo.UpdateUser(id, user)
}

func (s *userService) DeleteUser(id string) (*mongo.DeleteResult, error) {
	return s.repo.DeleteUser(id)
}
