package core

import (
	"context"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"strings"
	"time"
)

type (
	User struct {
		ID             uuid.UUID `gorm:"primary_key" json:"id"`
		Email          string    `gorm:"type:varchar;unique_index" json:"email"`
		Password       string    `gorm:"type:varchar" json:"password"`
		FirstName      string    `gorm:"type:varchar" json:"first_name"`
		SchedulerToken uuid.UUID `gorm:"unique_index" json:"scheduler_token"`
		SchedulerURL   *string   `gorm:"type:varchar" json:"scheduler_url"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}

	UserStore interface {
		Find(context.Context, uuid.UUID) (*User, error)

		FindByEmail(context.Context, string) (*User, error)

		Create(context.Context, *User) error
	}

	UserService interface {
		IssueToken(*User) (string, error)

		UserIdFromToken(token string) (*uuid.UUID, error)

		HashPassword(string) (string, error)

		ComparePassword(string, string) bool
	}
)

func (u *User) BeforeCreate(scope *gorm.Scope) error {
	u.ID = uuid.NewV4()
	u.SchedulerToken = uuid.NewV4()
	return nil
}

func (u *User) BeforeSave() error {
	u.Email = strings.ToLower(u.Email)
	return nil
}
