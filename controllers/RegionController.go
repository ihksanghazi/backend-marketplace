package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/backend-marketplace/model/web"
	"github.com/ihksanghazi/backend-marketplace/services"
)

type RegionController interface {
	GetAllProvince(c *gin.Context)
	GetCityByProvinceId(c *gin.Context)
}

type regionControllerImpl struct {
	service services.RegionService
}

func NewRegionController(service services.RegionService) RegionController {
	return &regionControllerImpl{
		service: service,
	}
}

func (r *regionControllerImpl) GetAllProvince(c *gin.Context) {
	result, err := r.service.GetAllProvince()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	response := web.BasicResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}

	c.JSON(200, response)

}

func (r *regionControllerImpl) GetCityByProvinceId(c *gin.Context) {
	provinceId := c.Param("provinceid")

	result, err := r.service.GetCityByProvinceId(provinceId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	response := web.BasicResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}

	c.JSON(200, response)
}
