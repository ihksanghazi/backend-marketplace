package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/backend-marketplace/model/web"
)

type StoreController interface{
	Create(c *gin.Context)
}

type storeControllerImpl struct{}

func NewStoreController() StoreController {
	return &storeControllerImpl{}
}

func (s *storeControllerImpl) Create(c *gin.Context) {
	var req web.CreateStoreRequest
	if err:=c.ShouldBindJSON(&req);err != nil {
		c.JSON(400,gin.H{"error":err.Error()})
		return
	}

	c.JSON(200,req)
}