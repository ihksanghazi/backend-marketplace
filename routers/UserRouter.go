package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/backend-marketplace/controllers"
)

func UserRouter(r *gin.RouterGroup){

	controller := controllers.NewUserController()

	r.POST("/register",controller.Register)
}