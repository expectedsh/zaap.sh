package core

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type (
	ApplicationState string

	ApplicationEnvironment map[string]string

	Application struct {
		ID          uuid.UUID              `gorm:"primary_key" json:"id"`
		Name        string                 `gorm:"type:varchar" json:"name"`
		Image       string                 `gorm:"type:varchar" json:"image"`
		Replicas    int                    `gorm:"type:integer" json:"replicas"`
		Environment ApplicationEnvironment `gorm:"type:json;not null" json:"environment"`
		State       ApplicationState       `gorm:"type:varchar" json:"state"`
		UserID      uuid.UUID              `json:"user_id"`
		CreatedAt   time.Time              `json:"created_at"`
		UpdatedAt   time.Time              `json:"updated_at"`

		User *User `json:"-"`
	}

	ApplicationStore interface {
		Find(context.Context, uuid.UUID) (*Application, error)

		List(context.Context, uuid.UUID) ([]*Application, error)

		Create(context.Context, *Application) error

		Delete(context.Context, uuid.UUID) error
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
	if a.Environment == nil {
		a.Environment = make(ApplicationEnvironment)
	}
	return nil
}

func (a ApplicationEnvironment) Value() (driver.Value, error) {
	b, err := json.Marshal(&a)
	if err != nil {
		return nil, err
	}
	return string(b), nil

}

func (a *ApplicationEnvironment) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), a)
}
