package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface{
	Register(c *gin.Context)
}

type UserControllerImpl struct{}

func NewUserController() UserController {
	return &UserControllerImpl{}
}

func (u *UserControllerImpl) Register(c *gin.Context) {
	c.JSON(http.StatusCreated,gin.H{"msg":"Register Successfully"})
}