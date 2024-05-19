package http

import "github.com/JerryJeager/mingle-backend/service/users"

type UserController struct {
	serv users.UserSv
}

func NewUserController(serv users.UserSv) *UserController {
	return &UserController{serv: serv}
}