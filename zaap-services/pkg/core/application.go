package core

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"time"
)

type (
	ApplicationState string

	Application struct {
		ID          uuid.UUID         `gorm:"primary_key" json:"id"`
		Name        string            `gorm:"type:varchar" json:"name"`
		Image       string            `gorm:"type:varchar" json:"image"`
		Replicas    int               `gorm:"type:integer" json:"replicas"`
		Environment map[string]string `gorm:"type:json" json:"environment"`
		State       ApplicationState  `gorm:"type:varchar" json:"state"`
		UserID      uuid.UUID         `json:"user_id"`
		CreatedAt   time.Time         `json:"created_at"`
		UpdatedAt   time.Time         `json:"updated_at"`

		User User `json:"-"`
	}

	ApplicationStore interface {
		List(context.Context, uuid.UUID) ([]*Application, error)
	}
)

const (
	ApplicationStateUnknown  ApplicationState = "unknown"
	ApplicationStateStarting                  = "starting"
	ApplicationStateRunning                   = "running"
	ApplicationStateStopping                  = "stopping"
	ApplicationStateStopped                   = "stopped"
	ApplicationStateCrashed                   = "crashed"
)
