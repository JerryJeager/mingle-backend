package http

import (
	"net/http"
	// "fmt"

	"github.com/JerryJeager/mingle-backend/service/users"
	"github.com/JerryJeager/mingle-backend/utils"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	serv users.UserSv
}

func NewUserController(serv users.UserSv) *UserController {
	return &UserController{serv: serv}
}

func (o *UserController) CreateUserWithGoogle(ctx *gin.Context){
	code := ctx.Query("code")
	// var pathUrl string = "/"

	// if ctx.Query("state") != "" {
	// 	pathUrl = ctx.Query("state")
	// }

	if code == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Authorization code not provided!"})
		return
	}

	tokenRes, err := utils.GetGoogleOauthToken(code)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	google_user, err := utils.GetGoogleUser(tokenRes.Access_token, tokenRes.Id_token)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	id, err := o.serv.CreateUserWithGoogle(ctx, google_user)

	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "user account already exists"})
	}

	ctx.JSON(http.StatusCreated, id)

	// ctx.Redirect(http.StatusTemporaryRedirect, fmt.Sprint(config.FrontEndOrigin, pathUrl))
}
