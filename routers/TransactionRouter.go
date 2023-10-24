package routers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/backend-marketplace/controllers"
	"github.com/ihksanghazi/backend-marketplace/services"
)

func TransactionRouter(r *gin.RouterGroup) {

	var ctx context.Context
	TrxService := services.NewTransactionService(ctx)
	CartService := services.NewCartService(ctx)
	controller := controllers.NewTransactionController(TrxService, CartService)

	r.GET("/ongkir/:id", controller.CekOngir)
	r.POST("/checkout/:id", controller.Checkout)
	r.GET("/user/:id", controller.GetByUserId)
	r.GET("/store/:id", controller.GetByStoreId)
}
