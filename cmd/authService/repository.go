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

// GetUserByID fetches a user by ID and tenantId
func (r *AuthRepository) GetUserByID(userID string, tenantId string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("id = ? AND tenant_id = ?", userID, tenantId).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

// GetUserByID fetches a user by UID and tenantId
func (r *AuthRepository) GetUserByUID(uId string, tenantId string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("uid = ? AND tenant_id = ?", uId, tenantId).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

// GetUserByEmail fetches a user by email and tenantId
func (r *AuthRepository) GetUserByEmail(email string, tenantId string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ? AND tenant_id = ?", email, tenantId).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

// UpdateUser updates an existing user's data with tenantId validation
func (r *AuthRepository) UpdateUser(user *models.User) (*models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Ensure tenant_id is part of the update condition
	if err := r.db.Where("id = ? AND tenant_id = ?", user.ID, user.TenantId).Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser deletes a user by ID and tenantId (soft delete)
func (r *AuthRepository) DeleteUser(userID string, tenantId string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if err := r.db.Where("id = ? AND tenant_id = ?", userID, tenantId).Delete(&models.User{}).Error; err != nil {
		return err
	}

	return nil
}
