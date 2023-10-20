package services

import (
	"context"

	"github.com/ihksanghazi/backend-marketplace/database"
	"github.com/ihksanghazi/backend-marketplace/model/domain"
	"github.com/ihksanghazi/backend-marketplace/model/web"
)

type RegionService interface {
	GetAllProvince() ([]web.Province, error)
	GetCityByProvinceId(provinceId string) ([]web.City, error)
}

type regionServiceImpl struct {
	ctx context.Context
}

func NewRegionService(ctx context.Context) RegionService {
	return &regionServiceImpl{
		ctx: ctx,
	}
}

func (r *regionServiceImpl) GetAllProvince() ([]web.Province, error) {
	var province domain.Province
	var response []web.Province
	err := database.DB.Model(province).WithContext(r.ctx).Find(&response).Error
	return response, err
}

func (r *regionServiceImpl) GetCityByProvinceId(provinceId string) ([]web.City, error) {
	var city domain.City
	var response []web.City
	err := database.DB.Model(city).WithContext(r.ctx).Where("province_id = ?", provinceId).Find(&response).Error
	return response, err
}
