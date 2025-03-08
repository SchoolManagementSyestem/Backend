package helpers

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Custom Claims Structure
type CustomClaims struct {
	UserID   uuid.UUID `json:"userId"`
	Role     string    `json:"role"`
	TenantID uuid.UUID `json:"tenantId"`
	jwt.RegisteredClaims
}

// GenerateAccessToken creates an access token
func GenerateAccessToken(userID uuid.UUID, role string, tenantID uuid.UUID) (string, error) {
	// Access Token Expiry
	accessTokenExpiryRaw := os.Getenv("ACCESS_TOKEN_EXPIRY")
	accessTokenExpiry, err := strconv.Atoi(accessTokenExpiryRaw)
	if err != nil {
		log.Fatal("Failed to parse ACCESS_TOKEN_EXPIRY")
	}

	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")

	claims := CustomClaims{
		UserID:   userID,
		Role:     role,
		TenantID: tenantID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(accessTokenExpiry))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(accessTokenSecret))
}

// GenerateRefreshToken creates a refresh token
func GenerateRefreshToken(userID uuid.UUID) (string, error) {

	// Refresh TOken Expiry
	refreshTokenExpiryRaw := os.Getenv("REFRESH_TOKEN_EXPIRY")
	refreshTokenExpiry, err := strconv.Atoi(refreshTokenExpiryRaw)
	if err != nil {
		log.Fatal("Failed to parse ACCESS_TOKEN_EXPIRY")
	}
	refreshTokenSecret := os.Getenv("REFRESH_TOKEN_SECRET")
	claims := jwt.RegisteredClaims{
		Subject:   userID.String(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(refreshTokenExpiry))),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(refreshTokenSecret))
}
