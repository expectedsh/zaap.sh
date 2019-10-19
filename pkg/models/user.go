package models

import (
  "github.com/dgrijalva/jwt-go"
  "github.com/jinzhu/gorm"
  uuid "github.com/satori/go.uuid"
  "time"
)

type User struct {
  ID        uuid.UUID `gorm:"type:uuid;primary_key"`
  Name      string    `gorm:"type:varchar;not null"`
  Email     string    `gorm:"type:varchar;not null;unique"`
  GithubID  int64     `gorm:"type:integer;unique"`
  GoogleID  string    `gorm:"type:varchar;unique"`
  UpdatedAt time.Time `gorm:"type:timestamptz;not null"`
  CreatedAt time.Time `gorm:"type:timestamptz;not null"`
}

func (u *User) BeforeCreate(scope *gorm.Scope) error {
  return scope.SetColumn("ID", uuid.NewV4())
}

func (u *User) NewToken() (string, error) {
  return jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
    ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
    Subject:   u.ID.String(),
  }).SigningString()
}
