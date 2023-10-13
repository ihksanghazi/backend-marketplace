package routers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/backend-marketplace/controllers"
	"github.com/ihksanghazi/backend-marketplace/middleware"
	"github.com/ihksanghazi/backend-marketplace/services"
)

func UserRouter(r *gin.RouterGroup){

	var ctx context.Context
	service := services.NewUserService(ctx)
	controller := controllers.NewUserController(service)
	middleware:= middleware.NewMiddleware(ctx)

	r.POST("/register",controller.Register)
	r.POST("/login",controller.Login)
	r.GET("/token",controller.GetToken)
	r.DELETE("/logout",controller.Logout)

	// must login
	r.Use(middleware.MustLogin())
	r.PUT("/:id",controller.Update)
	
}