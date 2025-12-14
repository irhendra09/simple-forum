package utils

import (
	"errors"
	"strings"
	"time"

	"donedev.com/simple-forum/internal/interfaces"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("secret")

// TokenSvc is the package-level token service. It can be set at startup.
var TokenSvc interfaces.TokenService

// JwtTokenService is a JWT-based implementation of interfaces.TokenService.
type JwtTokenService struct {
	key []byte
}

func NewJwtTokenService(key []byte) *JwtTokenService {
	return &JwtTokenService{key: key}
}

func (s *JwtTokenService) GenerateToken(userID int64) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Minute * 3).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.key)
}

func (s *JwtTokenService) GenerateRefreshToken(userID int64) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.key)
}

func (s *JwtTokenService) ParseToken(tokenStr string) (int64, error) {
	if strings.HasPrefix(tokenStr, "Bearer ") {
		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
	}
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return s.key, nil
	})
	if err != nil {
		return 0, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if v, ok := claims["user_id"].(float64); ok {
			return int64(v), nil
		}
		return 0, errors.New("invalid token claims")
	}
	return 0, errors.New("invalid token")
}

// package helpers - forward to TokenSvc when set, otherwise use default implementation
func GenerateToken(userID int64) (string, error) {
	if TokenSvc != nil {
		return TokenSvc.GenerateToken(userID)
	}
	svc := NewJwtTokenService(jwtKey)
	return svc.GenerateToken(userID)
}

func GenerateRefreshToken(userID int64) (string, error) {
	if TokenSvc != nil {
		return TokenSvc.GenerateRefreshToken(userID)
	}
	svc := NewJwtTokenService(jwtKey)
	return svc.GenerateRefreshToken(userID)
}

func ParseToken(token string) (int64, error) {
	if TokenSvc != nil {
		return TokenSvc.ParseToken(token)
	}
	svc := NewJwtTokenService(jwtKey)
	return svc.ParseToken(token)
}
