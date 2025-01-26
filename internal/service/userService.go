// internal/service/userService.go
package service

import (
	"fmt"

	"github.com/segmentio/kafka-go"
	"github.com/sumonskys/schoolManagementSystem/internal/kafkaInt"
	"github.com/sumonskys/schoolManagementSystem/internal/model"
	"github.com/sumonskys/schoolManagementSystem/internal/repository"
)

type UserService struct {
	userRepo repository.UserRepository
	producer *kafka.Writer
}

func NewUserService(userRepo repository.UserRepository, producer *kafka.Writer) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(user *model.User) error {
	err := s.userRepo.Save(user)
	if err != nil {
		return err
	}

	// Send message to Kafka after creating user
	message := fmt.Sprintf("User created: %s", user.Name)
	err = kafkaInt.SendMessage(s.producer, message)
	return err
}

func (s *UserService) GetUserByID(id int64) (*model.User, error) {
	return s.userRepo.FindByID(id)
}
