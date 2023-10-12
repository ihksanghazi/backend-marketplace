package controllers

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/backend-marketplace/model/web"
	"github.com/ihksanghazi/backend-marketplace/services"
	"github.com/jackc/pgx/v5/pgconn"
)

type UserController interface{
	Register(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
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
	if err != nil {
		var duplicateEntryError = &pgconn.PgError{Code: "23505"}
		if errors.As(err, &duplicateEntryError) {
			c.JSON(409,gin.H{"error":"email or phone number has already exist"})
			return
		}else{
			c.JSON(500,gin.H{"error":err.Error()})
			return
		}
	}
	
	response:= web.BasicResponse{
		Code: 201,
		Status: "Successful user registration",
		Data: res,
	}

	c.JSON(201,response)
}

func (u *userControllerImpl) Login(c *gin.Context) {
	
	var req web.LoginRequest
	if err:=c.ShouldBindJSON(&req);err!= nil {
		c.JSON(400,gin.H{"error":err.Error()})
		return
	}

	refreshToken,accessToken,err:=u.service.Login(req)
	if err != nil {
		if err.Error() == "wrong password"{
			c.JSON(401,gin.H{"error":err.Error()})
			return
		}else{
			c.JSON(500,gin.H{"error":err.Error()})
			return
		}
	}

	c.SetCookie("tkn_ck",refreshToken,int(time.Until(time.Now().Add(24 * time.Hour)).Seconds()),"/","localhost",false,true)

	c.JSON(200,gin.H{"your_access_token":accessToken})
}

func (u *userControllerImpl) Logout(c *gin.Context){
	c.SetCookie("tkn_ck","",-1,"/","localhost",false,true)
	c.JSON(200,gin.H{"msg":"berhasil logout"})
}