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
	Environment map[string]string

	Deployment struct {
		ID            uuid.UUID   `gorm:"primary_key" json:"id"`
		Image         string      `gorm:"type:varchar;not null" json:"image"`
		Replicas      int         `gorm:"type:integer;not null" json:"replicas"`
		Environment   Environment `gorm:"type:json;not null" json:"environment"`
		ApplicationID uuid.UUID   `json:"application_id"`
		CreatedAt     time.Time   `json:"created_at"`
		UpdatedAt     time.Time   `json:"updated_at"`

		Application *Application `json:"-"`
	}

	DeploymentStore interface {
		List(context.Context, uuid.UUID) (*[]*Deployment, error)

		Create(context.Context, *Deployment) error
	}
)

func (d *Deployment) BeforeCreate(scope *gorm.Scope) error {
	if d.ID == uuid.Nil {
		d.ID = uuid.NewV4()
	}
	if d.Environment == nil {
		d.Environment = make(Environment)
	}
	return nil
}

func (d Environment) Value() (driver.Value, error) {
	b, err := json.Marshal(&d)
	if err != nil {
		return nil, err
	}
	return string(b), nil

}

func (d *Environment) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), d)
}
