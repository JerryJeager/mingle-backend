package http

import (
	"net/http"
	"os"

	// "fmt"

	"github.com/JerryJeager/mingle-backend/models"
	"github.com/JerryJeager/mingle-backend/service/users"
	"github.com/JerryJeager/mingle-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		// ctx.JSON(http.StatusOK, gin.H{"id": id, "token": token})
		handleFrontendRedirect(ctx, id, token)
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
	handleFrontendRedirect(ctx, id, token)
}

func (o *UserController) CreateUser(ctx *gin.Context) {
	var user models.CreateUserReq

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "email and password is required"})
		return
	}

	id, username, err := o.serv.CreateUser(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "user with email already exists"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": id, "username": username})
}

func (o *UserController) GetUser(ctx *gin.Context) {
	id, ok := ctx.Get("user_id")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "faliled to get user id from context"})
		return
	}
	user, err := o.serv.GetUser(ctx, uuid.MustParse(id.(string)))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "user with email already exists"})
		return
	}

	ctx.JSON(http.StatusOK, *user)
}

func (o *UserController) CreateToken(ctx *gin.Context) {
	var user models.CreateUserReq

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "email and password is required"})
		return
	}

	id, token, err := o.serv.CreateToken(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid email or password"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": id, "token": token})
	ctx.SetCookie("user_id", id, 86400, "/", "localhost", false, true)
	ctx.SetCookie("access_token", token, 86400, "/", "localhost", false, true)
}

func handleFrontendRedirect(ctx *gin.Context, id, token string) {
	environment := os.Getenv("ENVIRONMENT")
	var path string
	var redirect string
	if environment == "development" {
		path = "localhost"
		redirect = "http://localhost:3000/dashboard"
	} else {
		path = "we-mingle.vercel.app"
		redirect = "https://we-mingle.vercel.app/dashboard"
	}
	ctx.SetCookie("user_id", id, 86400, "/", path, false, false)
	ctx.SetCookie("access_token", token, 86400, "/", path, false, false)

	ctx.Redirect(http.StatusTemporaryRedirect, redirect)
}
