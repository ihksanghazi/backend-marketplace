package controllers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/backend-marketplace/model/web"
	"github.com/ihksanghazi/backend-marketplace/services"
	"github.com/ihksanghazi/backend-marketplace/utils"
)

type TransactionController interface {
	CekOngir(c *gin.Context)
}

type tranctionControllerImpl struct {
	service services.TransactionService
}

func NewTransactionController(service services.TransactionService) TransactionController {
	return &tranctionControllerImpl{
		service: service,
	}
}

func (t *tranctionControllerImpl) CekOngir(c *gin.Context) {
	cartId := c.Param("id")
	expedition := c.DefaultQuery("expedition", "jne")

	if expedition != "jne" && expedition != "pos" && expedition != "tiki" {
		c.JSON(400, gin.H{"error": "must be jne, pos, or tiki"})
		return
	}

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

	result, err := t.service.CekOngkir(cartId, claims.ID, expedition)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	response := web.BasicResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}

	c.JSON(200, response)
}
