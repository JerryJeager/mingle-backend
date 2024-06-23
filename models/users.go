package models

import (
	"html"
	"strings"
	"time"

	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID  `json:"id" gorm:"unique;type:uuid"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
	Username  string     `json:"username" gorm:"unique"`
	Email     string     `json:"email" gorm:"primary_key"`
	Password  string     `json:"password"`
	Gender    string     `json:"gender"`
	Picture   string     `json:"picture"`
	AuthType  string     `json:"auth_type" gorm:"primary_key"`
}

const (
	Google = "google"
	Normal = "normal"
)

type CreateUserReq struct {
	Email    string `json:"email" gorm:"unique" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (user *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	user.Email = html.EscapeString(strings.TrimSpace(user.Email))

	return nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
