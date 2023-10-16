package routers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/backend-marketplace/controllers"
	"github.com/ihksanghazi/backend-marketplace/middleware"
	"github.com/ihksanghazi/backend-marketplace/services"
)

func ProductRouter(r *gin.RouterGroup) {

	var ctx context.Context
	service := services.NewProductService(ctx)
	controller := controllers.NewProductController(service)
	middleware:=middleware.NewMiddleware(ctx)

	r.Use(middleware.MustLogin())
	r.POST("/create", controller.Create)
	r.PUT("/:id", controller.Update)
	r.DELETE("/:id", controller.Delete)
	r.GET("/find", controller.Find)
	r.GET("/:id", controller.Get)

}
