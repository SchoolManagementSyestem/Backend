package helpers

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
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

// ParseToken verifies and extracts claims from an access token
func ParseToken(tokenString string) (*CustomClaims, error) {
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(accessTokenSecret), nil
	})

	if err != nil {
		return nil, err
	}

	// Extract and return claims if valid
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// ExtractTokenFromHeader gets JWT from the Authorization header
func ExtractTokenFromHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("missing authorization header")
	}

	// Authorization: Bearer <TOKEN>
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("invalid authorization header format")
	}

	return parts[1], nil
}

// VerifyTokenFromRequest extracts and verifies token from HTTP request
func VerifyTokenFromRequest(r *http.Request) (*CustomClaims, error) {
	tokenString, err := ExtractTokenFromHeader(r)
	if err != nil {
		return nil, err
	}

	return ParseToken(tokenString)
}
