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
	Checkout(c *gin.Context)
}

type tranctionControllerImpl struct {
	TrxService  services.TransactionService
	CartService services.CartService
}

func NewTransactionController(TrxService services.TransactionService, CartService services.CartService) TransactionController {
	return &tranctionControllerImpl{
		TrxService:  TrxService,
		CartService: CartService,
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

	result, err := t.TrxService.CekOngkir(cartId, claims.ID, expedition)
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

func (t *tranctionControllerImpl) Checkout(c *gin.Context) {
	cartId := c.Param("id")
	payment := c.Query("payment")

	var req web.CheckoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if payment != "bca" && payment != "bni" && payment != "bri" {
		c.JSON(400, gin.H{"error": "Must be bca, Bni, Or Bri"})
		return
	}

	result, err := t.TrxService.Checkout(cartId, req, payment)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// delete cart
	if err := t.CartService.DeleteCart(cartId); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, result)

}
