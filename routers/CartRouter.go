package routers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/backend-marketplace/controllers"
	"github.com/ihksanghazi/backend-marketplace/services"
)

func CartRouter(r *gin.RouterGroup) {

	var ctx context.Context
	cartService := services.NewCartService(ctx)
	controller := controllers.NewCartController(cartService)

	r.POST("/add/:product_id", controller.Add)
	r.DELETE("/:id", controller.DeleteCart)
	r.PUT("/item/:id", controller.UpdateItem)
	r.DELETE("/item/:id", controller.DeleteItem)
	r.GET("/get", controller.Get)
}
