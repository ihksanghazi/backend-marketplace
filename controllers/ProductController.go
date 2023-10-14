package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/backend-marketplace/model/web"
)

type ProductController interface{
	Create(c *gin.Context)
}

type productControllerImpl struct{}

func NewProductController() ProductController {
	return &productControllerImpl{}
}

func (p *productControllerImpl) Create(c *gin.Context) {
	var req web.CreateProductRequest
	if err:=c.ShouldBindJSON(&req);err != nil {
		c.JSON(400,gin.H{"error":err.Error()})
		return
	}

	c.JSON(200,req)
}