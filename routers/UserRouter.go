package routers

import "github.com/gin-gonic/gin"

func UserRouter(r *gin.RouterGroup){
	r.POST("/register",func(c *gin.Context) {
		c.JSON(200,gin.H{"msg":"berhasil"})
	})
}