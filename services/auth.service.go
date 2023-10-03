package services

import (
	"errors"
	api_errors "scylla-grpc-adapter/internal/errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtManager struct {
	SecretKey			string
	TokenDuration time.Duration
}

type UserClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

func Register(user *User) error {
	err := CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func Login(username, password string) (*User, error) {
	user, err := GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	match, err := ComparePasswordAndHash(password, user.Password)
	if err != nil {
		return nil, err
	}

	if !match {
		return nil, errors.New(api_errors.ErrInvalidPassword)
	}

	return user, nil
}

func NewJwtManager(secretKey string, tokenDuration time.Duration) *JwtManager {
	return &JwtManager{secretKey, tokenDuration}
}

func (manager *JwtManager) Generate(user *User) (string, error) {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.TokenDuration).Unix(),
		},
		Username: user.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.SecretKey))
}

func (manager *JwtManager) Verify(token string) (*UserClaims, error) {
	claims := &UserClaims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(manager.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !tkn.Valid {
		return nil, errors.New(api_errors.ErrInvalidToken)
	}

	return claims, nil
}
