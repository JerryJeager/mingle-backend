package users

import (
	"context"

	"github.com/JerryJeager/mingle-backend/models"
	"github.com/JerryJeager/mingle-backend/utils"
	"github.com/google/uuid"
)

type UserSv interface {
	CreateUserWithGoogle(context context.Context, user *utils.GoogleUserResult) (string, error)
	LoginUserWithGoogle(ctx context.Context, email string) (string, string, error)
}

type UserServ struct {
	repo UserStore
}

func NewUserService(repo UserStore) *UserServ {
	return &UserServ{repo: repo}
}

func (o *UserServ) CreateUserWithGoogle(context context.Context, user *utils.GoogleUserResult) (string, error) {
	id := uuid.New()

	err := o.repo.CreateUserWithGoogle(context, user, id)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (o *UserServ) LoginUserWithGoogle(ctx context.Context, email string) (string, string, error) {
	id, err := o.repo.GetUserID(ctx, email)
	if err != nil {
		return "", "", err
	}

	user := models.User{
		ID: id,
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		return "", "", err
	}

	return id.String(), token, nil
}
