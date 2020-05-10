package core

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"github.com/satori/go.uuid"
	"regexp"
	"time"
)

type (
	Environment map[string]string

	Deployment struct {
		ID            uuid.UUID      `gorm:"primary_key" json:"id"`
		Image         string         `gorm:"type:varchar;not null" json:"image"`
		Replicas      int            `gorm:"type:integer;not null" json:"replicas"`
		Environment   Environment    `gorm:"type:json;not null" json:"environment"`
		Roles         pq.StringArray `gorm:"type:varchar[]" json:"roles"`
		ApplicationID uuid.UUID      `json:"application_id"`
		CreatedAt     time.Time      `json:"created_at"`

		Application *Application `json:"-"`
	}

	DeploymentStore interface {
		Find(context.Context, uuid.UUID) (*Deployment, error)

		ListByApplication(context.Context, uuid.UUID) (*[]*Deployment, error)

		Create(context.Context, *Deployment) error
	}
)

var DeploymentImageRegex = regexp.MustCompile("(?m)^(?:.+/)?([^:]+)(?::.+)?$")

func (d *Deployment) BeforeCreate(scope *gorm.Scope) error {
	if d.ID == uuid.Nil {
		d.ID = uuid.NewV4()
	}
	return nil
}

func (d *Deployment) BeforeSave(scope *gorm.Scope) error {
	if d.Environment == nil {
		d.Environment = make(Environment)
	}
	if d.Roles == nil {
		d.Roles = pq.StringArray{}
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
