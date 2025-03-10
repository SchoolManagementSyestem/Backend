package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	// Generate a hashed password with a cost factor of 10
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckPasswordHash(password, hashedPassword string) error {
	// Compare the password with the hashed password
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
