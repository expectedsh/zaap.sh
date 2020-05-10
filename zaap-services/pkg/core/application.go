package core

import (
	"context"
	"fmt"
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/runnerpb"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/protocol"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"github.com/satori/go.uuid"
	"regexp"
	"strings"
	"time"
)

type (
	ApplicationStatus string

	Application struct {
		ID                  uuid.UUID         `gorm:"primary_key" json:"id"`
		Name                string            `gorm:"type:varchar;not null" json:"name"`
		Status              ApplicationStatus `gorm:"type:varchar;not null" json:"status"`
		DefaultDomain       string            `gorm:"type:varchar;unique;not null" json:"default_domain"`
		Domains             pq.StringArray    `gorm:"type:varchar[];not null" json:"domains"`
		UserID              uuid.UUID         `json:"user_id"`
		RunnerID            uuid.UUID         `json:"runner_id"`
		CurrentDeploymentID uuid.UUID         `json:"current_deployment_id"`
		CreatedAt           time.Time         `json:"created_at"`
		UpdatedAt           time.Time         `json:"updated_at"`

		User              *User         `json:"-"`
		Runner            *Runner       `json:"-"`
		CurrentDeployment *Deployment   `json:"-"`
		Deployments       []*Deployment `json:"deployments,omitempty"`
	}

	ApplicationStore interface {
		Find(context.Context, uuid.UUID) (*Application, error)

		FindWithRunner(context.Context, uuid.UUID) (*Application, error)

		ListByUser(context.Context, uuid.UUID) (*[]Application, error)

		ListByRunner(context.Context, uuid.UUID) (*[]Application, error)

		Create(context.Context, *Application) error

		Update(context.Context, *Application) error

		Delete(context.Context, uuid.UUID) error
	}

	ApplicationService interface {
		Deploy(*Application) error

		NotifyCreated(*Application) error

		NotifyUpdated(*Application) error

		NotifyDeleted(*Application) error

		NotifyStatusChanged(*Application) error
	}
)

const (
	ApplicationStatusUnknown   ApplicationStatus = "unknown"
	ApplicationStatusDeploying                   = "deploying"
	ApplicationStatusRunning                     = "running"
	ApplicationStatusCrashed                     = "crashed"
	ApplicationStatusFailed                      = "failed"
)

var ApplicationNameRegex = regexp.MustCompile("(?m)^[a-z]([-a-z0-9]*[a-z0-9])?$")

func (a *Application) BeforeCreate(scope *gorm.Scope) error {
	if a.ID == uuid.Nil {
		a.ID = uuid.NewV4()
	}
	if a.DefaultDomain == "" {
		a.DefaultDomain = fmt.Sprintf("%v.gtw.zaap.sh", strings.ReplaceAll(a.ID.String(), "-", ""))
	}
	return nil
}

func (a *Application) BeforeSave(scope *gorm.Scope) error {
	if a.Domains == nil {
		a.Domains = pq.StringArray{}
	}
	return nil
}

func ApplicationStatusFromRunner(status runnerpb.ApplicationStatus) ApplicationStatus {
	switch status {
	case runnerpb.ApplicationStatus_DEPLOYING:
		return ApplicationStatusDeploying
	case runnerpb.ApplicationStatus_RUNNING:
		return ApplicationStatusRunning
	case runnerpb.ApplicationStatus_CRASHED:
		return ApplicationStatusCrashed
	case runnerpb.ApplicationStatus_FAILED:
		return ApplicationStatusFailed
	default:
		return ApplicationStatusUnknown
	}
}

func (s ApplicationStatus) ToMessagingFormat() protocol.ApplicationStatus {
	switch s {
	case ApplicationStatusDeploying:
		return protocol.ApplicationStatus_A_DEPLOYING
	case ApplicationStatusRunning:
		return protocol.ApplicationStatus_A_RUNNING
	case ApplicationStatusCrashed:
		return protocol.ApplicationStatus_A_CRASHED
	case ApplicationStatusFailed:
		return protocol.ApplicationStatus_A_FAILED
	default:
		return protocol.ApplicationStatus_A_UNKNOWN
	}
}
