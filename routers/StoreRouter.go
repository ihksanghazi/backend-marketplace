package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/backend-marketplace/controllers"
)

func StoreRouter(r *gin.RouterGroup) {

	controller:=controllers.NewStoreController()

	r.POST("/create",controller.Create)
	
}