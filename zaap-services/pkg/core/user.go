package core

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"strings"
	"time"
)

type (
	User struct {
		ID             uuid.UUID `json:"id"`
		Email          string    `json:"email"`
		Password       string    `json:"password"`
		FirstName      string    `json:"first_name"`
		SchedulerToken uuid.UUID `json:"scheduler_token"`
		SchedulerURL   string    `json:"scheduler_url"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}

	UserStore interface {
		Find(context.Context, uuid.UUID) (*User, error)

		Create(context.Context, *User) error
	}

	UserService interface {
		IssueToken(*User) (string, error)
	}
)

func (u *User) BeforeSave() error {
	u.Email = strings.ToLower(u.Email)
	return nil
}
