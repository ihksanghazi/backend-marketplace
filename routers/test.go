package routers

import (
	"os"

	"github.com/gin-gonic/gin"
)

func TestRouter(r *gin.RouterGroup){
	r.GET("/",func(c *gin.Context) {
		c.JSON(200,gin.H{"message":os.Getenv("TEST")})
	})
}