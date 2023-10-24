package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/backend-marketplace/controllers"
)

func ReviewRouter(r *gin.RouterGroup) {

	controller := controllers.NewReviewController()

	r.POST("/create/:productId", controller.Create)
}
