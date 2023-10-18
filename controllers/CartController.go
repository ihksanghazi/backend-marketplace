package controllers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/backend-marketplace/model/web"
	"github.com/ihksanghazi/backend-marketplace/services"
	"github.com/ihksanghazi/backend-marketplace/utils"
	"gorm.io/gorm"
)

type CartController interface {
	Add(c *gin.Context)
	DeleteCart(c *gin.Context)
}

type cartControllerImpl struct {
	service services.CartService
}

func NewCartController(service services.CartService) CartController {
	return &cartControllerImpl{
		service: service,
	}
}

func (ca *cartControllerImpl) Add(c *gin.Context) {
	productID := c.Param("product_id")
	amount := c.DefaultQuery("qty", "1")
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

	if err := ca.service.Add(productID, amount, claims.ID); err != nil {
		if err.Error() == "you are the selller of this item" {
			c.JSON(409, gin.H{"error": err.Error()})
			return
		} else if err == gorm.ErrRecordNotFound || err.Error() == "product is not available" {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	result, err := ca.service.Get(claims.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	response := web.BasicResponse{
		Code:   201,
		Status: "Success Add Product With Id '" + productID + "'" + " To Your Cart",
		Data:   result,
	}

	c.JSON(201, response)
}

func (ca *cartControllerImpl) DeleteCart(c *gin.Context) {
	id := c.Param("id")

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

	if err := ca.service.DeleteCart(id); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	result, err := ca.service.Get(claims.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	response := web.BasicResponse{
		Code:   200,
		Status: "Success Delete Cart With Id '" + id + "'",
		Data:   result,
	}

	c.JSON(200, response)
}
