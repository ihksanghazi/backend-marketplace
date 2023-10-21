package routers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/backend-marketplace/controllers"
	"github.com/ihksanghazi/backend-marketplace/services"
)

func TransactionRouter(r *gin.RouterGroup) {

	var ctx context.Context
	service := services.NewTransactionService(ctx)
	controller := controllers.NewTransactionController(service)

	r.GET("/ongkir/:id", controller.CekOngir)
	r.POST("checkout/:id", controller.Checkout)
}
