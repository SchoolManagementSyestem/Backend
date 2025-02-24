package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// AuthService handles authentication logic
type AuthService struct {
	secretKey string
	users     map[string]string // Mock user storage (username -> hashed password)
}

// NewAuthService initializes the authentication service
func NewAuthService(secretKey string) *AuthService {
	return &AuthService{
		secretKey: secretKey,
		users: map[string]string{
			"admin": hashPassword("secret"), // Mocked user
		},
	}
}

// Authenticate checks user credentials and returns a token
func (as *AuthService) Authenticate(username, password string) (string, error) {
	hashedPassword, exists := as.users[username]
	if !exists {
		return "", errors.New("user not found")
	}

	// Verify password
	if !checkPasswordHash(password, hashedPassword) {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT token
	token, err := as.generateToken(username)
	if err != nil {
		return "", err
	}

	return token, nil
}

// generateToken creates a JWT token for the user
func (as *AuthService) generateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 2).Unix(), // Token expires in 2 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(as.secretKey))
}

// hashPassword hashes a plain text password
func hashPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed)
}

// checkPasswordHash compares a hashed password with a plaintext password
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
