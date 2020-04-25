package store

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type deploymentStore struct {
	db *gorm.DB
}

func NewDeploymentStore(db *gorm.DB) core.DeploymentStore {
	return &deploymentStore{db}
}

func (s deploymentStore) Find(ctx context.Context, id uuid.UUID) (*core.Deployment, error) {
	deployment := new(core.Deployment)
	if err := s.db.First(deployment, "id = ?", id.String()).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return deployment, nil
}

func (s deploymentStore) ListByApplication(ctx context.Context, applicationId uuid.UUID) (*[]*core.Deployment, error) {
	deployments := new([]*core.Deployment)
	if err := s.db.Find(deployments, "application_id = ?", applicationId.String()).Error; err != nil {
		return nil, err
	}
	return deployments, nil
}

func (s deploymentStore) Create(ctx context.Context, deployment *core.Deployment) error {
	return s.db.Create(deployment).Error
}
