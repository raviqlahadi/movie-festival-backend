package services

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthService struct {
	SecretKey string
}

func NewAuthService(secretKey string) *AuthService {
	return &AuthService{SecretKey: secretKey}
}

func (s *AuthService) GenerateToken(userID uint, isAdmin bool) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"is_admin": isAdmin,
		"exp":      time.Now().Add(72 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.SecretKey))
}
