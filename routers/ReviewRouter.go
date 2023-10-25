package routers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/backend-marketplace/controllers"
	"github.com/ihksanghazi/backend-marketplace/services"
)

func ReviewRouter(r *gin.RouterGroup) {

	var ctx context.Context
	service := services.NewReviewService(ctx)
	controller := controllers.NewReviewController(service)

	r.POST("/create/:productId", controller.Create)
	r.GET("/get/:productId", controller.Get)
	r.PUT("/:reviewId", controller.Update)
}
