package controllers

import (
	"net/http"
	"strconv"

	"github.com/alireza-mf/go-ecommerce/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func ProvideUserController(u services.UserService) UserController {
	return UserController{UserService: u}
}

func (u *UserController) GetUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err := u.UserService.FindById(uint(userId))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": http.StatusText(500)})
		return
	}
	if &user == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": http.StatusText(404)})
		return
	}

	c.JSON(http.StatusOK, &user)
}
