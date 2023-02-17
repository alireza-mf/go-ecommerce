package controllers

import (
	"net/http"
	"strings"

	"github.com/alireza-mf/go-ecommerce/models"
	"github.com/alireza-mf/go-ecommerce/services"
	"github.com/alireza-mf/go-ecommerce/util/jwt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	UserService services.UserService
}

func ProvideUserController(u services.UserService) UserController {
	return UserController{UserService: u}
}

func (u *UserController) GetUser(c *gin.Context) {
	userId := c.Param("user_id")

	user, err := u.UserService.FindById(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": http.StatusText(500)})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": http.StatusText(404)})
		return
	}

	c.JSON(http.StatusOK, &user)
}

func (u *UserController) RegisterUser(c *gin.Context) {
	var body models.UserInputForm
	err := c.ShouldBindBodyWith(&body, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var response *models.User
	response, err = u.UserService.CreateUser(&body)
	if err != nil {
		if strings.Contains(err.Error(), "CreateUser::email_is_existed") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email is existed."})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, &response)
}

func (u *UserController) LoginUser(c *gin.Context) {
	var body models.UserLoginForm
	err := c.ShouldBindBodyWith(&body, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err := u.UserService.UserRepository.FindByEmail(body.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": http.StatusText(404)})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	accessToken, err := jwt.GenerateJWT(user.Email, user.UserId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}
