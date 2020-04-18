package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
)

type userService struct {
	secretKey string
}

func NewUserService(secretKey string) core.UserService {
	return &userService{secretKey}
}

func (s *userService) IssueToken(user *core.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID.String(),
	})
	return token.SignedString(s.secretKey)
}
