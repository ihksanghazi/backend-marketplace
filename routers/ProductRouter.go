package routers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/backend-marketplace/controllers"
	"github.com/ihksanghazi/backend-marketplace/services"
)

func ProductRouter(r *gin.RouterGroup){

	var ctx context.Context
	service:=services.NewProductService(ctx)
	controller:=controllers.NewProductController(service)
	r.POST("/create",controller.Create)
	r.PUT("/:id",controller.Update)
	r.DELETE("/:id",controller.Delete)
	
}