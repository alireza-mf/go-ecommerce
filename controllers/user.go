package controllers

import (
	"net/http"
	"strings"

	"github.com/alireza-mf/go-ecommerce/models"
	"github.com/alireza-mf/go-ecommerce/services"
	"github.com/alireza-mf/go-ecommerce/util"
	"github.com/alireza-mf/go-ecommerce/util/jwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	UserService services.UserService
}

func ProvideUserController(u services.UserService) UserController {
	return UserController{UserService: u}
}

func (u *UserController) RegisterUser(c *gin.Context) {
	input := c.MustGet("input").(models.RegisterUserInput)

	response, err := u.UserService.CreateUser(&input.Body)
	if err != nil {
		if strings.Contains(err.Error(), "CreateUser::email_is_existed") {
			util.ResponseError(c, http.StatusBadRequest, "Email is existed.")
		} else {
			util.ResponseError(c, http.StatusInternalServerError)
		}
		return
	}

	util.ResponseSuccess(c, &response)
}

func (u *UserController) LoginUser(c *gin.Context) {
	input := c.MustGet("input").(models.LoginUserInput)

	user, err := u.UserService.UserRepository.FindByEmail(input.Body.Email)
	if err != nil {
		util.ResponseError(c, http.StatusInternalServerError)
		return
	}
	if user == nil {
		util.ResponseError(c, http.StatusBadRequest, "Incorrect email or password.")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Body.Password))
	if err != nil {
		util.ResponseError(c, http.StatusBadRequest, "Incorrect email or password")
		return
	}

	accessToken, err := jwt.GenerateJWT(user.Email, user.UserId)
	if err != nil {
		util.ResponseError(c, http.StatusInternalServerError)
		return
	}

	util.ResponseSuccess(c, struct{ access_token string }{access_token: accessToken})
}

func (u *UserController) GetUser(c *gin.Context) {
	input := c.MustGet("input").(models.GetUserInput)

	user, err := u.UserService.FindById(input.Params.UserId)

	if err != nil {
		util.ResponseError(c, http.StatusInternalServerError)
		return
	}
	if user == nil {
		util.ResponseError(c, http.StatusNotFound)
		return
	}

	util.ResponseSuccess(c, &user)
}
