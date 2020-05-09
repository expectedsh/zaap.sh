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

func (s applicationStore) Find(ctx context.Context, id uuid.UUID) (*core.Application, error) {
	application := new(core.Application)
	if err := s.db.Preload("User").First(application, "id = ?", id.String()).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return application, nil
}

func (s applicationStore) FindWithRunner(ctx context.Context, id uuid.UUID) (*core.Application, error) {
	application := new(core.Application)
	if err := s.db.Preload("User").Preload("Runner").First(application, "id = ?", id.String()).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return application, nil
}

func (s applicationStore) ListByUser(ctx context.Context, userId uuid.UUID) (*[]core.Application, error) {
	applications := new([]core.Application)
	if err := s.db.Find(applications, "user_id = ?", userId).Error; err != nil {
		return nil, err
	}
	return applications, nil
}

func (s applicationStore) ListByRunner(ctx context.Context, runnerId uuid.UUID) (*[]core.Application, error) {
	applications := new([]core.Application)
	if err := s.db.Find(applications, "runner_id = ?", runnerId).Error; err != nil {
		return nil, err
	}
	return applications, nil
}

func (s applicationStore) Create(ctx context.Context, application *core.Application) error {
	return s.db.Create(application).Error
}

func (s applicationStore) Update(ctx context.Context, application *core.Application) error {
	return s.db.Save(application).Error
}

func (s applicationStore) Delete(ctx context.Context, id uuid.UUID) error {
	return s.db.Delete(&core.Application{}, "id = ?", id).Error
}
