package users

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
	Username  string     `json:"username" gorm:"unique"`
	Email     string     `json:"email" gorm:"unique"`
	Password  string     `json:"password"`
	Gender    string     `json:"gender"`
	Picture   string     `json:"picture"`
}
