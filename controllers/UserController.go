package controllers

import (
	"errors"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/backend-marketplace/model/web"
	"github.com/ihksanghazi/backend-marketplace/services"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type UserController interface{
	Register(c *gin.Context)
	Login(c *gin.Context)
	GetToken(c *gin.Context)
	Logout(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Find(c *gin.Context)
	GetUser(c *gin.Context)
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
			c.JSON(409,gin.H{"error":err.Error()})
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
		}else if err == gorm.ErrRecordNotFound{
			c.JSON(404,gin.H{"error":err.Error()})
			return
		}else{
			c.JSON(500,gin.H{"error":err.Error()})
			return
		}
	}

	c.SetCookie("tkn_ck",refreshToken,int(time.Until(time.Now().Add(24 * time.Hour)).Seconds()),"/","localhost",false,true)

	c.JSON(200,gin.H{"your_access_token":accessToken})
}

func (u *userControllerImpl) GetToken(c *gin.Context){
	// get token
	refreshToken,err:=c.Cookie("tkn_ck")
	if err != nil || refreshToken == "" {
		c.JSON(401,gin.H{"error":"Unauthorize"})
		return
	}

	accessToken,err:=u.service.GetToken(refreshToken)
	if err != nil {
		if err == gorm.ErrRecordNotFound{
			c.JSON(401,gin.H{"error":"Unauthorize"})
			return
		}else{
			c.JSON(500,gin.H{"error":err.Error()})
			return
		}
	}

	c.JSON(200,gin.H{"your_access_token":accessToken})
}

func (u *userControllerImpl) Logout(c *gin.Context){
	c.SetCookie("tkn_ck","",-1,"/","localhost",false,true)
	c.JSON(200,gin.H{"msg":"Successfull logout"})
}

func (u *userControllerImpl) Update(c *gin.Context){
	// get url
	id := c.Param("id")

	var req web.UpdateRequest
	if err:= c.ShouldBindJSON(&req);err!= nil {
		c.JSON(400,gin.H{"error":err.Error()})
		return
	}

	res,err:=u.service.Update(id,req)
	if err!= nil {
		if err == gorm.ErrRecordNotFound{
			c.JSON(404,gin.H{"error":err.Error()})
			return
		}else{
			c.JSON(500,gin.H{"error":err.Error()})
			return
		}
	}

	response:=web.BasicResponse{
		Code: 200,
		Status: "Successfull Update User With Id '"+id+"'",
		Data: res,
	}

	c.JSON(200,response)
}

func (u *userControllerImpl) Delete(c *gin.Context){
	id := c.Param("id")

	if err:=u.service.Delete(id);err != nil {
		if err == gorm.ErrRecordNotFound{
			c.JSON(404,gin.H{"error":err.Error()})
			return
		}else{
			c.JSON(500,gin.H{"error":err.Error()})
			return
		}
	}

	c.JSON(200,gin.H{"msg":"Success delete user with id '"+id+"'"})
}

func (u *userControllerImpl) Find(c *gin.Context){
	page:=c.DefaultQuery("page","1")
	limit:=c.DefaultQuery("limit","5")
	search:=c.DefaultQuery("search","")

	page1,err:=strconv.Atoi(page)
	if err != nil {
		c.JSON(400,gin.H{"error":err.Error()})
		return
	}
	limit1,err:=strconv.Atoi(limit)
	if err != nil {
		c.JSON(400,gin.H{"error":err.Error()})
		return
	}

	result,totalPage,err:=u.service.Find(page1,limit1,search)
	if err != nil {
		c.JSON(500,gin.H{"error":err.Error()})
		return
	}

	response:= web.Pagination{
		Code: 200,
		Status: "OK",
		CurrentPage: page,
		TotalPage: totalPage,
		Data: result,
	}

	c.JSON(200,response)
}

func (u *userControllerImpl) GetUser(c *gin.Context){
	id:=c.Param("id")

	result,err:=u.service.GetUser(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound{
			c.JSON(404,gin.H{"error":err.Error()})
			return
		}else{
			c.JSON(500,gin.H{"error":err.Error()})
			return
		}
	}

	response:= web.BasicResponse{
		Code: 200,
		Status: "OK",
		Data: result,
	}

	c.JSON(200,response)
}