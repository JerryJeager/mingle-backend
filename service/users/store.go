package users

import (
	"context"

	"github.com/JerryJeager/mingle-backend/config"
	"github.com/JerryJeager/mingle-backend/models/users"
	"github.com/JerryJeager/mingle-backend/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserStore interface {
	CreateUserWithGoogle(context context.Context, user *utils.GoogleUserResult, userID uuid.UUID) error
}

type UserRepo struct {
	client *gorm.DB
}

func NewUserRepo(client *gorm.DB) *UserRepo {	
	return &UserRepo{client: client}
}

func (o *UserRepo) CreateUserWithGoogle(context context.Context, user *utils.GoogleUserResult, userID uuid.UUID) error {
	newUser := users.User{
		ID: userID,
		Email: user.Email,
		Picture: user.Picture,
		Username: utils.RandomUserName(),
	}
	
	result := config.Session.Create(&newUser).WithContext(context)

	if result.Error != nil{
		return result.Error
	}

	return nil
}