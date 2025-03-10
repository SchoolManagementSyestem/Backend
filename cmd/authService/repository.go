package main

import (
	"errors"
	"schoolManagementSystem/internal/db/models"
	"sync"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
	mu sync.Mutex
}

// NewAuthRepository initializes a new repository instance
func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

// Save stores a new user in the database
func (r *AuthRepository) Save(user *models.User) (*models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByID fetches a user by ID
func (r *AuthRepository) GetUserByID(userID string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("id = ?", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

// GetUserByEmail fetches a user by email
func (r *AuthRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

// UpdateUser updates an existing user's data
func (r *AuthRepository) UpdateUser(user *models.User) (*models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser deletes a user by ID (soft delete)
func (r *AuthRepository) DeleteUser(userID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if err := r.db.Where("id = ?", userID).Delete(&models.User{}).Error; err != nil {
		return err
	}

	return nil
}
