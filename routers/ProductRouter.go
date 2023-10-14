package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/backend-marketplace/controllers"
)

func ProductRouter(r *gin.RouterGroup){

	controller:=controllers.NewProductController()
	r.POST("/create",controller.Create)
	
}