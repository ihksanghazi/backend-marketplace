package controllers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/backend-marketplace/model/web"
	"github.com/ihksanghazi/backend-marketplace/services"
	"github.com/ihksanghazi/backend-marketplace/utils"
)

type ReviewController interface {
	Create(c *gin.Context)
}

type reviewControllerImpl struct {
	service services.ReviewService
}

func NewReviewController(service services.ReviewService) ReviewController {
	return &reviewControllerImpl{
		service: service,
	}
}

func (r *reviewControllerImpl) Create(c *gin.Context) {
	productId := c.Param("productId")

	refreshToken, err := c.Cookie("tkn_ck")
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	claims, err := utils.ParsingToken(refreshToken, os.Getenv("REFRESH_TOKEN"))
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	var req web.CreateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := r.service.Create(claims.ID, productId, req); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"msg": "Success Create Product Review"})
}
