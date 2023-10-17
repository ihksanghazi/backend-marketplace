package routers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/backend-marketplace/controllers"
	"github.com/ihksanghazi/backend-marketplace/services"
)

func CartRouter(r *gin.RouterGroup) {

	var ctx context.Context
	service := services.NewCartService(ctx)
	controller := controllers.NewCartController(service)

	r.POST("/add/:product_id", controller.Add)
}
