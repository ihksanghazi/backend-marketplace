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

type ProductController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Find(c *gin.Context)
	Get(c *gin.Context)
}

type productControllerImpl struct {
	service services.ProductService
}

func NewProductController(service services.ProductService) ProductController {
	return &productControllerImpl{
		service: service,
	}
}

func (p *productControllerImpl) Create(c *gin.Context) {
	// get refresh token
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

	var req web.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := p.service.Create(claims.ID, req); err != nil {
		if err.Error() == "you don't have a shop" {
			c.JSON(409, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(201, gin.H{"msg": "Success Create Product"})
}

func (p *productControllerImpl) Update(c *gin.Context) {
	id := c.Param("id")

	var req web.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := p.service.Update(id, req)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	response := web.BasicResponse{
		Code:   200,
		Status: "Success Update Product With Id '" + id + "'",
		Data:   result,
	}

	c.JSON(200, response)
}

func (p *productControllerImpl) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := p.service.Delete(id); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(200, gin.H{"msg": "Success Delete Product With Id '" + id + "'"})
}

func (p *productControllerImpl) Find(c *gin.Context) {
	search := c.DefaultQuery("search", "")
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	page1, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	limit1, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, totalPage, err := p.service.Find(search, page1, limit1)
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

func (p *productControllerImpl) Get(c *gin.Context) {
	id := c.Param("id")

	result, err := p.service.Get(id)
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
		Status: "OK",
		Data:   result,
	}

	c.JSON(200, response)
}
