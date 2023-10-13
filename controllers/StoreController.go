package controllers

import (
	"errors"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/backend-marketplace/model/web"
	"github.com/ihksanghazi/backend-marketplace/services"
	"github.com/ihksanghazi/backend-marketplace/utils"
	"github.com/jackc/pgx/v5/pgconn"
)

type StoreController interface{
	Create(c *gin.Context)
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