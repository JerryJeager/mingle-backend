package users

import (
	"context"

	"github.com/JerryJeager/mingle-backend/config"
	"github.com/JerryJeager/mingle-backend/models"
	"github.com/JerryJeager/mingle-backend/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserStore interface {
	CreateUserWithGoogle(context context.Context, user *utils.GoogleUserResult, userID uuid.UUID) error
	GetUserID(ctx context.Context, email string) (uuid.UUID, error)
	CreateUser(ctx context.Context, user *models.User) error
	GetUserPassword(ctx context.Context, userEmail string) (string, error)
}

type UserRepo struct {
	client *gorm.DB
}

func NewUserRepo(client *gorm.DB) *UserRepo {
	return &UserRepo{client: client}
}

func (o *UserRepo) CreateUserWithGoogle(context context.Context, user *utils.GoogleUserResult, userID uuid.UUID) error {
	newUser := models.User{
		ID:       userID,
		Email:    user.Email,
		Picture:  user.Picture,
		Username: utils.RandomUserName(),
	}

	result := config.Session.Create(&newUser).WithContext(context)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (o *UserRepo) GetUserID(ctx context.Context, email string) (uuid.UUID, error) {
	var user models.User
	query := config.Session.First(&user, "email = ?", email).WithContext(ctx)
	if query.Error != nil {
		return uuid.UUID{}, query.Error
	}
	return user.ID, nil
}

func (o *UserRepo) CreateUser(ctx context.Context, user *models.User) error {

	query := config.Session.Create(&user).WithContext(ctx)

	if query.Error != nil {
		return query.Error
	}

	return nil
}

func (o *UserRepo) GetUserPassword(ctx context.Context, userEmail string) (string, error) {
	var userModel models.User

	if err := config.Session.First(&userModel, "email = ?", userEmail).WithContext(ctx).Error; err != nil {
		return "", err
	}

	return userModel.Password, nil
}
