package controllers

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/backend-marketplace/model/web"
	"github.com/ihksanghazi/backend-marketplace/services"
	"github.com/ihksanghazi/backend-marketplace/utils"
	"gorm.io/gorm"
)

type ReviewController interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
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

func (r *reviewControllerImpl) Get(c *gin.Context) {
	productId := c.Param("productId")
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	page_int, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	limit_int, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, totalPage, err := r.service.Get(productId, page_int, limit_int)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	response := web.Pagination{
		Code:        200,
		Status:      "OK",
		CurrentPage: page,
		TotalPage:   totalPage,
		Data:        result,
	}

	c.JSON(200, response)
}

func (r *reviewControllerImpl) Update(c *gin.Context) {
	reviewId := c.Param("reviewId")

	var req web.UpdateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := r.service.Update(reviewId, req)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
	}

	response := web.BasicResponse{
		Code:   200,
		Status: "Success Update Review With Id '" + reviewId + "'",
		Data:   result,
	}

	c.JSON(200, response)
}

func (r *reviewControllerImpl) Delete(c *gin.Context) {
	reviewId := c.Param("reviewId")
	c.JSON(200, gin.H{"id": reviewId})
}
