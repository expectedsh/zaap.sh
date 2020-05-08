package store

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type runnerStore struct {
	db *gorm.DB
}

func NewRunnerStore(db *gorm.DB) core.RunnerStore {
	return &runnerStore{db}
}

func (s runnerStore) Find(ctx context.Context, id uuid.UUID) (*core.Runner, error) {
	runner := new(core.Runner)
	if err := s.db.First(runner, "id = ?", id.String()).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return runner, nil
}

func (s runnerStore) List(ctx context.Context) (*[]core.Runner, error) {
	runners := new([]core.Runner)
	if err := s.db.Find(runners).Error; err != nil {
		return nil, err
	}
	return runners, nil
}

func (s runnerStore) ListByUser(ctx context.Context, userId uuid.UUID) (*[]core.Runner, error) {
	runners := new([]core.Runner)
	if err := s.db.Find(runners, "user_id = ?", userId).Error; err != nil {
		return nil, err
	}
	return runners, nil
}

func (s runnerStore) Create(ctx context.Context, runner *core.Runner) error {
	return s.db.Create(runner).Error
}

func (s runnerStore) Update(ctx context.Context, runner *core.Runner) error {
	return s.db.Save(runner).Error
}

func (s runnerStore) Delete(ctx context.Context, id uuid.UUID) error {
	return s.db.Delete(&core.Runner{}, "id = ?", id).Error
}
