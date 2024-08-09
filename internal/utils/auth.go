package utils

import (
	"Houses/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	UserID   uint           `json:"user_id"`
	UserType model.UserType `json:"user_type"`
	jwt.RegisteredClaims
}

type AuthManager struct {
	secretKey []byte
}

func NewAuthManager(secretKey []byte) *AuthManager {
	return &AuthManager{secretKey: secretKey}
}

//////////////////////////////////////////////////////////////////////////////////////////////////

const tokenLifetime = time.Hour * 24

func (a *AuthManager) GenerateToken(id uint, utype model.UserType) (string, error) {
	claims := Claims{
		UserID:   id,
		UserType: utype,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenLifetime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(a.secretKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func (a *AuthManager) ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return a.secretKey, nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
