package routers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/backend-marketplace/controllers"
	"github.com/ihksanghazi/backend-marketplace/services"
)

func RegionRouter(r *gin.RouterGroup) {

	var ctx context.Context
	service := services.NewRegionService(ctx)
	controller := controllers.NewRegionController(service)

	r.GET("/province", controller.GetAllProvince)
	r.GET("/city/:provinceid", controller.GetCityByProvinceId)
}
