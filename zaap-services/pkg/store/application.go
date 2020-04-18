package store

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type applicationStore struct {
	db *gorm.DB
}

func NewApplicationStore(db *gorm.DB) core.ApplicationStore {
	return &applicationStore{db}
}

func (s applicationStore) List(ctx context.Context, userId uuid.UUID) ([]*core.Application, error) {
	applications := new([]*core.Application)
	if err := s.db.Find(applications, "user_id = ?", userId).Error; err != nil {
		return nil, err
	}
	return *applications, nil
}
