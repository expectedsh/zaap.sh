package service

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"github.com/satori/go.uuid"
)

type userService struct {
	secretKey string
}

func NewUserService(secretKey string) core.UserService {
	return &userService{secretKey}
}

func (s userService) IssueToken(user *core.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID.String(),
	})
	return token.SignedString([]byte(s.secretKey))
}

func (s userService) UserIdFromToken(rawToken string) (*uuid.UUID, error) {
	token, err := jwt.Parse(rawToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := uuid.FromString(claims["user_id"].(string))
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func (s userService) HashPassword(password string) (string, error) {
	hash := sha512.New()
	if _, err := hash.Write([]byte(password)); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func (s userService) ComparePassword(hashedPassword string, password string) bool {
	hash, err := s.HashPassword(password)
	if err != nil {
		return false
	}
	return hashedPassword == hash
}
