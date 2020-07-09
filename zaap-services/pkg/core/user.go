package core

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"strings"
	"time"
)

type (
	User struct {
		ID        uuid.UUID `gorm:"primary_key" json:"id"`
		Email     string    `gorm:"type:varchar;unique_index;not null" json:"email"`
		Password  string    `gorm:"type:varchar;not null" json:"-"`
		FirstName string    `gorm:"type:varchar;not null" json:"first_name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`

		Applications []*Application `json:"-"`
	}

	UserStore interface {
		Find(context.Context, uuid.UUID) (*User, error)

		FindByEmail(context.Context, string) (*User, error)

		Create(context.Context, *User) error

		Update(context.Context, *User) error
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
	return nil
}

func (u *User) BeforeSave() error {
	u.Email = strings.ToLower(u.Email)
	return nil
}
