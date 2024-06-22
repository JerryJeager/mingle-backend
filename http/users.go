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

func (o *UserController) CreateUserWithGoogle(ctx *gin.Context) {
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
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail 1", "message": err.Error()})
		return
	}

	google_user, err := utils.GetGoogleUser(tokenRes.Access_token, tokenRes.Id_token)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail 2", "message": err.Error()})
		return
	}

	_, err = o.serv.CreateUserWithGoogle(ctx, google_user)

	if err != nil {
		//TODO: CHECK IF USER ALREADY EXISTS BEFORE SIGN IN:
		id, token, err := o.serv.LoginUserWithGoogle(ctx, google_user.Email)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"id": id, "token": token})
		return
	}

	// ctx.JSON(http.StatusCreated, id)
	//login after signup with google
	id, token, err := o.serv.LoginUserWithGoogle(ctx, google_user.Email)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	// ctx.JSON(http.StatusOK, gin.H{"id": id, "token": token})
	ctx.SetCookie("user_id", id, 86400, "/", "localhost", false, true)
	ctx.SetCookie("access_token", token, 86400, "/", "https://we-mingle.vercel.app", false, true)

	ctx.Redirect(http.StatusTemporaryRedirect, "https://we-mingle.vercel.app/dashboard")
}
