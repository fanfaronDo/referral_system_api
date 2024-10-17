package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/fanfaronDo/referral_system_api/pkg/model"
	"github.com/fanfaronDo/referral_system_api/pkg/storage"
	"time"
)

const (
	tokenExpireTime = time.Hour * 12
	signedKey       = "KwkdkfewowelosklKalosdk"
	salt            = "HAUWFGiwgbkwsGeHGeuh"
)

type claims struct {
	jwt.StandardClaims
	UserId uint
}

type Auth struct {
	storage *storage.Storage
}

func NewAuth(storage *storage.Storage) *Auth {
	return &Auth{storage}
}

func (s *Auth) CreateUser(user *model.User) error {
	pass := s.generateHashForPassword(user.Password)
	user.Password = pass
	return s.storage.CreateUser(user)
}

func (s *Auth) GenerateToken(username, password string) (string, error) {
	user, err := s.storage.AuthStorage.GetUser(username, s.generateHashForPassword(password))

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenExpireTime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signedKey))
}

func (a *Auth) ParseToken(accessToken string) (uint, error) {
	token, err := jwt.ParseWithClaims(accessToken, &claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrUnexpectedSigningMethod
		}

		return []byte(signedKey), nil
	})
	if err != nil {
		return 0, err
	}

	c, ok := token.Claims.(*claims)
	if !ok || !token.Valid {
		return 0, ErrInvalidToken
	}

	return c.UserId, nil
}

func (s *Auth) generateHashForPassword(pass string) string {
	hash := sha1.New()
	hash.Write([]byte(pass))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
