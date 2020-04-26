package core

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/protocol"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"regexp"
	"time"
)

type (
	RunnerType string

	RunnerStatus string

	Runner struct {
		ID          uuid.UUID      `gorm:"primary_key" json:"id"`
		Name        string         `gorm:"type:varchar;not null" json:"name"`
		Description *string        `gorm:"type:varchar" json:"description"`
		Type        RunnerType     `gorm:"type:varchar;not null" json:"type"`
		ExternalIps pq.StringArray `gorm:"type:varchar[];not null" json:"external_ips"`
		Status      RunnerStatus   `gorm:"type:varchar;not null" json:"status"`
		Url         string         `gorm:"type:varchar;not null" json:"url"`
		Token       string         `gorm:"type:varchar;not null" json:"token"`
		UserID      uuid.UUID      `json:"user_id"`
		CreatedAt   time.Time      `json:"created_at"`
		UpdatedAt   time.Time      `json:"updated_at"`

		User *User `json:"-"`
	}

	RunnerStore interface {
		Find(context.Context, uuid.UUID) (*Runner, error)

		ListByUser(context.Context, uuid.UUID) (*[]Runner, error)

		Create(context.Context, *Runner) error

		Update(context.Context, *Runner) error

		Delete(context.Context, uuid.UUID) error
	}

	RunnerService interface {
		NewConnection(*Runner) (protocol.RunnerClient, *grpc.ClientConn, error)
	}
)

const (
	RunnerTypeDockerSwarm RunnerType   = "docker_swarm"
	RunnerTypeKubernetes               = "kubernetes"
	RunnerTypeUnknown                  = "unknown"
	RunnerStatusUnknown   RunnerStatus = "unknown"
	RunnerStatusOnline                 = "online"
	RunnerStatusOffline                = "offline"
)

func (r *Runner) BeforeCreate(scope *gorm.Scope) error {
	r.ID = uuid.NewV4()
	return nil
}

func (r *Runner) BeforeSave(scope *gorm.Scope) error {
	if r.ExternalIps == nil {
		r.ExternalIps = pq.StringArray{}
	}
	return nil
}

func (r *Runner) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(
			&r.Name,
			validation.Required,
			validation.Length(3, 50),
			validation.Match(regexp.MustCompile("(?m)^[-a-zA-Z0-9]+$")).
				Error("should only contain letters, numbers, and dashes"),
		),
		validation.Field(&r.Description, validation.Length(0, 255)),
		validation.Field(&r.Url, validation.Required),
		validation.Field(&r.Token, validation.Required, validation.Length(8, 255)),
	)
}
