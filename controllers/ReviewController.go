package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/backend-marketplace/model/web"
)

type ReviewController interface {
	Create(c *gin.Context)
}

type reviewControllerImpl struct{}

func NewReviewController() ReviewController {
	return &reviewControllerImpl{}
}

func (r *reviewControllerImpl) Create(c *gin.Context) {
	productId := c.Param("productId")
	var req web.CreateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"id": productId, "req": req})
}
