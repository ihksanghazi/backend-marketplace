package middleware

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/backend-marketplace/database"
	"github.com/ihksanghazi/backend-marketplace/model/domain"
	"github.com/ihksanghazi/backend-marketplace/utils"
	"gorm.io/gorm"
)

type Middleware interface{
	MustLogin() gin.HandlerFunc
	MustAdmin() gin.HandlerFunc
}

type middlewareImpl struct{
	ctx context.Context
}

func NewMiddleware(ctx context.Context) Middleware{
	return &middlewareImpl{
		ctx:ctx,
	}
}

func (m *middlewareImpl) MustLogin() gin.HandlerFunc{
	return func(c *gin.Context){
		// get token
		accessToken:=c.GetHeader("Access-Token")
		// cek token
		if accessToken == "" {
			c.AbortWithStatusJSON(401,gin.H{"error":"Unauthorized"})
			return
		}
		
		// parse access Token
		_,err:=utils.ParsingToken(accessToken, os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			c.AbortWithStatusJSON(401,gin.H{"error":"Unauthorized"})
			return
		}
		
		c.Next()
	}
}

func (m *middlewareImpl) MustAdmin() gin.HandlerFunc{
	return func(c *gin.Context) {
		// get token
		refreshToken,err:=c.Cookie("tkn_ck")
		// cek token
		if err != nil || refreshToken == "" {
			c.AbortWithStatusJSON(401,gin.H{"error":"Unauthorized"})
			return
		}
		// cek refresh token
		claims,err:=utils.ParsingToken(refreshToken,os.Getenv("REFRESH_TOKEN"))
		if err != nil {
			c.AbortWithStatusJSON(401,gin.H{"error":"Unauthorized"})
			return
		}
		// get user by id
		var user domain.User
		if err:=database.DB.Model(user).WithContext(m.ctx).Where("id = ?",claims.ID).First(&user).Error; err != nil {
			switch err {
				case gorm.ErrRecordNotFound:
					c.AbortWithStatusJSON(401,gin.H{"error":"Unauthorized"})
					return
				default:
					c.AbortWithStatusJSON(500,gin.H{"error":err.Error()})
					return
			}
		}

		// cek admin
		if user.Role != "admin" {
			c.AbortWithStatusJSON(401,gin.H{"error":"Not admin"})
			return
		}

		c.Next()
	}
}