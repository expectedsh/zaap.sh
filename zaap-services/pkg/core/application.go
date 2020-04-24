package core

import (
	"context"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type (
	ApplicationState string

	Application struct {
		ID                  uuid.UUID        `gorm:"primary_key" json:"id"`
		Name                string           `gorm:"type:varchar" json:"name"`
		State               ApplicationState `gorm:"type:varchar" json:"state"`
		UserID              uuid.UUID        `json:"user_id"`
		CurrentDeploymentID uuid.UUID        `json:"current_deployment_id"`
		CreatedAt           time.Time        `json:"created_at"`
		UpdatedAt           time.Time        `json:"updated_at"`

		User              *User         `json:"-"`
		CurrentDeployment *Deployment   `json:"-"`
		Deployments       []*Deployment `json:"deployments,omitempty"`
	}

	ApplicationStore interface {
		Find(context.Context, uuid.UUID) (*Application, error)

		List(context.Context, uuid.UUID) (*[]Application, error)

		Create(context.Context, *Application) error

		Update(context.Context, *Application) error

		Delete(context.Context, uuid.UUID) error
	}

	ApplicationService interface {
		Deploy(*Application) error

		NotifyDeletion(*Application) error
	}
)

const (
	ApplicationStateUnknown   ApplicationState = "unknown"
	ApplicationStateDeploying                  = "deploying"
	ApplicationStateRunning                    = "running"
	ApplicationStateCrashed                    = "crashed"
)

func (a *Application) BeforeCreate(scope *gorm.Scope) error {
	a.ID = uuid.NewV4()
	return nil
}
