package store

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type userStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) core.UserStore {
	return &userStore{db}
}

func (s *userStore) Find(ctx context.Context, id uuid.UUID) (*core.User, error) {
	user := new(core.User)
	if err := s.db.Find(user, "id = ?", id.String()).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userStore) Create(ctx context.Context, user *core.User) error {
	return s.db.Create(user).Error
}
