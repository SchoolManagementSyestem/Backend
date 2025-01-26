// internal/repository/userRepo.go
package repository

import "github.com/sumonskys/schoolManagementSystem/internal/model"

type UserRepository interface {
	FindByID(id int64) (*model.User, error)
	Save(user *model.User) error
}

// In-memory implementation
type InMemoryUserRepository struct {
	users map[int64]*model.User
}

func NewUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[int64]*model.User),
	}
}

func (repo *InMemoryUserRepository) FindByID(id int64) (*model.User, error) {
	user, exists := repo.users[id]
	if !exists {
		return nil, nil
	}
	return user, nil
}

func (repo *InMemoryUserRepository) Save(user *model.User) error {
	repo.users[user.ID] = user
	return nil
}
