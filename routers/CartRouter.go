package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/backend-marketplace/controllers"
)

func CartRouter(r *gin.RouterGroup) {

	controller := controllers.NewCartController()

	r.POST("/add/:product_id", controller.Add)
}
