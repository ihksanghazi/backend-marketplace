package routers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/backend-marketplace/middleware"
)

func TestRouter(r *gin.RouterGroup){

	var ctx context.Context
	middleware := middleware.NewMiddleware(ctx)

	// must login
	r.Use(middleware.MustLogin())
	r.GET("/middleware",func(c *gin.Context) {
		c.JSON(200,gin.H{"message":"Berhasil Akses Login"})
	})

	r.Use(middleware.MustAdmin())
	r.GET("/admin",func(c *gin.Context) {
		c.JSON(200,gin.H{"message":"Berhasil Akses Admin"})
	})
}