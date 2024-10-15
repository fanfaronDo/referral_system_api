package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/fanfaronDo/referral_system_api/internal/entry"
	"github.com/fanfaronDo/referral_system_api/internal/storage"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	tokenExpireTime = time.Hour * 12
	signedKey       = "KwkdkfewowelosklKalosdk"
)

type claims struct {
	jwt.StandardClaims
	UserId   uint
	Username string
}

type Auth struct {
	storage storage.Auth
}

func NewAuth(storage storage.Auth) *Auth {
	return &Auth{storage}
}

func (s *Auth) CreateUser(user *entry.User) error {
	pass, err := s.generateHashForPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = pass
	return s.storage.CreateUser(user)
}

func (s *Auth) GenerateToken(username, password string) (string, error) {
	hash, err := s.generateHashForPassword(password)
	if err != nil {
		return "", err
	}

	user, err := s.storage.GetUser(username, hash)
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenExpireTime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
		user.Username,
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

func (s *Auth) generateHashForPassword(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return string(hash), err
}
