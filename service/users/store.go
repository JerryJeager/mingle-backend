package users

import (
	"gorm.io/gorm"
)

type UserStore interface {
	
}

type UserRepo struct {
	client *gorm.DB
}

func NewUserRepo(client *gorm.DB) *UserRepo {	
	return &UserRepo{client: client}
}