package controllers

import (
	"errors"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/backend-marketplace/model/web"
	"github.com/ihksanghazi/backend-marketplace/services"
	"github.com/ihksanghazi/backend-marketplace/utils"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type StoreController interface{
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Find(c *gin.Context)
	Get(c *gin.Context)
}

type storeControllerImpl struct{
	service services.StoreService
}

func NewStoreController(service services.StoreService) StoreController {
	return &storeControllerImpl{
		service:service,
	}
}

func (s *storeControllerImpl) Create(c *gin.Context) {
	refreshToken,err:=c.Cookie("tkn_ck")
	if err!= nil {
		c.JSON(401,gin.H{"error":"Unauthorized"})
		return
	}

	claims,err:=utils.ParsingToken(refreshToken, os.Getenv("REFRESH_TOKEN"))
	if err!= nil {
		c.JSON(401,gin.H{"error":"Unauthorized"})
		return
	}

	var req web.CreateStoreRequest
	if err:=c.ShouldBindJSON(&req);err != nil {
		c.JSON(400,gin.H{"error":err.Error()})
		return
	}

	result,err:=s.service.Create(claims.ID,req)
	if err != nil {
		var duplicateEntryError = &pgconn.PgError{Code: "23505"}
		if errors.As(err, &duplicateEntryError) {
			c.JSON(409,gin.H{"error":"Prohibited from opening any more shops"})
			return
		}else{
			c.JSON(500,gin.H{"error":err.Error()})
			return
		}
	}

	response:=web.BasicResponse{
		Code: 201,
		Status: "Success Create Store",
		Data: result,
	}

	c.JSON(201,response)
}

func (s *storeControllerImpl) Update(c *gin.Context) {
	id:=c.Param("id")

	var req web.UpdateStoreRequest
	if err:=c.ShouldBindJSON(&req);err != nil {
		c.JSON(400,gin.H{"error":err.Error()})
		return
	}

	result,err:=s.service.Update(id,req)
	if err != nil {
		if err == gorm.ErrRecordNotFound{
			c.JSON(404,gin.H{"error":err.Error()})
			return
		}else{
			c.JSON(500,gin.H{"error":err.Error()})
			return
		}
	}

	response:=web.BasicResponse{
		Code:200,
		Status: "Success Update Store With Id '"+id+"'",
		Data: result,
	}

	c.JSON(200,response)
}

func (s *storeControllerImpl) Delete(c *gin.Context){
	id:=c.Param("id")

	if err:=s.service.Delete(id); err != nil {
		if err == gorm.ErrRecordNotFound{
			c.JSON(404,gin.H{"error":err.Error()})
			return
		}else{
			c.JSON(500,gin.H{"error":err.Error()})
			return
		}
	}

	c.JSON(200,gin.H{"msg":"Success Delete Store with id '"+id+"'"})
}

func (s *storeControllerImpl) Find(c *gin.Context) {
	page:=c.DefaultQuery("page","1")
	limit:=c.DefaultQuery("limit","10")
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

	result,totalPage,err:=s.service.Find(page1,limit1,search)
	if err != nil {
		c.JSON(500,gin.H{"error":err.Error()})
		return
	}

	pagination:=web.Pagination{
		Code: 200,
		Status: "OK",
		CurrentPage: page,
		TotalPage: totalPage,
		Data: result,
	}

	c.JSON(200,pagination)
}

func (s *storeControllerImpl) Get(c *gin.Context) {
	id:=c.Param("id")

	result,err:=s.service.Get(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound{
			c.JSON(404,gin.H{"error":err.Error()})
			return
		}else{
			c.JSON(500,gin.H{"error":err.Error()})
			return
		}
	}

	response:=web.BasicResponse{
		Code:200,
		Status: "OK",
		Data: result,
	}

	c.JSON(200,response)
}