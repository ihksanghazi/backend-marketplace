package routers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/backend-marketplace/controllers"
	"github.com/ihksanghazi/backend-marketplace/services"
)

func StoreRouter(r *gin.RouterGroup) {

	var ctx context.Context
	service:=services.NewStoreService(ctx)
	controller:=controllers.NewStoreController(service)

	r.POST("/create",controller.Create)
	
}