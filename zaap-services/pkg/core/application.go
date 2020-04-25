package core

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"regexp"
	"time"
)

type (
	ApplicationStatus string

	Application struct {
		ID                  uuid.UUID         `gorm:"primary_key" json:"id"`
		Name                string            `gorm:"type:varchar;not null" json:"name"`
		Status              ApplicationStatus `gorm:"type:varchar;not null" json:"state"`
		UserID              uuid.UUID         `json:"user_id"`
		RunnerID            uuid.UUID         `json:"runner_id"`
		CurrentDeploymentID uuid.UUID         `json:"current_deployment_id"`
		CreatedAt           time.Time         `json:"created_at"`
		UpdatedAt           time.Time         `json:"updated_at"`

		User              *User         `json:"-"`
		Runner            *Runner       `json:"runner"`
		CurrentDeployment *Deployment   `json:"-"`
		Deployments       []*Deployment `json:"deployments,omitempty"`
	}

	ApplicationStore interface {
		Find(context.Context, uuid.UUID) (*Application, error)

		ListByUser(context.Context, uuid.UUID) (*[]Application, error)

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
	ApplicationStatusUnknown   ApplicationStatus = "unknown"
	ApplicationStatusDeploying                   = "deploying"
	ApplicationStatusRunning                     = "running"
	ApplicationStatusCrashed                     = "crashed"
)

var ApplicationNameRegex = regexp.MustCompile("(?m)^[-a-zA-Z0-9]+$")

func (a *Application) BeforeCreate(scope *gorm.Scope) error {
	a.ID = uuid.NewV4()
	return nil
}

func (a *Application) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(
			&a.Name,
			validation.Required,
			validation.Length(3, 50),
			validation.Match(regexp.MustCompile("(?m)^[-a-zA-Z0-9]+$")).
				Error("should only contain letters, numbers, and dashes"),
		),
	)
}
