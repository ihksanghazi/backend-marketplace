package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/backend-marketplace/controllers"
)

func TransactionRouter(r *gin.RouterGroup) {

	controller := controllers.NewTransactionController()

	r.GET("/ongkir", controller.CekOngir)
}
