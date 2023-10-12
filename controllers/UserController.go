package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/backend-marketplace/model/web"
	"github.com/ihksanghazi/backend-marketplace/services"
	"github.com/jackc/pgx/v5/pgconn"
)

type UserController interface{
	Register(c *gin.Context)
}

type userControllerImpl struct{
	service services.UserService
}

func NewUserController(service services.UserService) UserController {
	return &userControllerImpl{
		service: service,
	}
}

func (u *userControllerImpl) Register(c *gin.Context) {
	var req web.RegisterRequest
	if err :=c.ShouldBindJSON(&req);err!= nil {
		c.JSON(400,gin.H{"error":err.Error()})
		return
	}

	res,err:=u.service.Register(req)
	// jika error duplikat
	var duplicateEntryError = &pgconn.PgError{Code: "23505"}
	if errors.As(err, &duplicateEntryError) {
		c.JSON(409,gin.H{"error":"email or phone number has already exist"})
		return
	}
	// jika error lainnya
	if err!= nil {
		c.JSON(500,gin.H{"error":err.Error()})
		return
	}
	
	response:= web.BasicResponse{
		Code: 201,
		Status: "Successful user registration",
		Data: res,
	}

	c.JSON(201,response)
}