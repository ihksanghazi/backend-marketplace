package controllers

import "github.com/gin-gonic/gin"

type CartController interface {
	Add(c *gin.Context)
}

type cartControllerImpl struct{}

func NewCartController() CartController {
	return &cartControllerImpl{}
}

func (ca *cartControllerImpl) Add(c *gin.Context) {
	productID := c.Param("product_id")
	amount := c.DefaultQuery("qty", "1")

	c.JSON(200, gin.H{"product_id": productID, "amount": amount})
}
