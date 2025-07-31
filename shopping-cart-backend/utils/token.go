package utils

import (
	"fmt" // Used for fmt.Errorf
	"time"

	"github.com/dgrijalva/jwt-go"
)

// IMPORTANT: In a production application, this secret key should be loaded from
// environment variables (e.g., os.Getenv("JWT_SECRET")) and be a strong, random string.
var jwtSecret = []byte("c0Lg68NuDQFHAYvB/yFi8cgzyR1Q4NcIASupV834tJw9yOud+dNlMMPi5T3WpT90dj8Rvt1gNffJLGztnZS30A==")

// GenerateToken creates a new JWT for a given user ID.
func GenerateToken(userID uint) (string, error) {
	// Define the claims (payload) for the token
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	// Create the token with the signing method and claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ValidateToken parses and validates a JWT string, returning the user ID if valid.
func ValidateToken(tokenString string) (uint, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"]) // fmt.Errorf is used here
		}
		return jwtSecret, nil
	})

	// Check if the token is valid and extract claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := uint(claims["user_id"].(float64)) // JWT claims parse numbers as float64
		return userID, nil
	} else {
		return 0, err
	}
}
