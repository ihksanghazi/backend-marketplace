package services

import (
	"context"
	"time"

	"github.com/ihksanghazi/backend-marketplace/database"
	"github.com/ihksanghazi/backend-marketplace/model/domain"
	"github.com/ihksanghazi/backend-marketplace/model/web"
	"gorm.io/gorm"
)

type StoreService interface {
	Create(userId string, req web.CreateStoreRequest) error
	Update(storeId string, req web.UpdateStoreRequest) (web.UpdateStoreRequest, error)
	Delete(storeId string) error
	Find(page int, limit int, search string) (result []web.FindStoreResponse, totalPage int, err error)
	Get(storeId string) (web.GetStoreResponse, error)
	Report(storeId string, startDate time.Time, endDate time.Time) (web.StoreReport, error)
}

type storeServiceImpl struct {
	ctx context.Context
}

func NewStoreService(ctx context.Context) StoreService {
	return &storeServiceImpl{
		ctx: ctx,
	}
}

func (s *storeServiceImpl) Create(userId string, req web.CreateStoreRequest) error {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		var user domain.User
		if err := tx.Model(user).WithContext(s.ctx).Where("id = ?", userId).First(&user).Error; err != nil {
			return err
		}
		var store domain.Store
		store.UserId = user.Id
		store.StoreName = req.StoreName
		store.Description = req.Description
		store.Category = req.Category
		store.ImageUrl = req.ImageUrl
		store.Address = user.Address
		store.CityId = user.CityId

		if err := tx.Model(store).WithContext(s.ctx).Create(&store).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (s *storeServiceImpl) Update(storeId string, req web.UpdateStoreRequest) (web.UpdateStoreRequest, error) {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		var store domain.Store
		store.StoreName = req.StoreName
		store.Description = req.Description
		store.Category = req.Category
		store.ImageUrl = req.ImageUrl
		store.Address = req.Address
		store.CityId = req.CityId
		if err := tx.Model(store).WithContext(s.ctx).Where("id = ?", storeId).Updates(store).First(&req).Error; err != nil {
			return err
		}
		return nil
	})
	return req, err
}

func (s *storeServiceImpl) Delete(storeId string) error {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		var store domain.Store
		if err := tx.Model(store).WithContext(s.ctx).Where("id = ?", storeId).First(&store).Error; err != nil {
			return err
		}
		if err := tx.Model(store).WithContext(s.ctx).Where("id = ?", storeId).Delete(&store).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (s *storeServiceImpl) Find(page int, limit int, search string) (result []web.FindStoreResponse, totalPage int, err error) {
	var store domain.Store
	var totalData int64
	offset := (page - 1) * limit
	var response []web.FindStoreResponse
	Err := database.DB.Model(store).WithContext(s.ctx).Where("store_name ILIKE ? OR category ILIKE ?", "%"+search+"%", "%"+search+"%").Count(&totalData).Limit(limit).Offset(offset).Find(&response).Error
	TotalPage := (int(totalData) + limit - 1) / limit
	return response, TotalPage, Err
}

func (s *storeServiceImpl) Get(storeId string) (web.GetStoreResponse, error) {
	var store domain.Store
	var response web.GetStoreResponse
	err := database.DB.Model(store).WithContext(s.ctx).Where("id = ?", storeId).Preload("Region").Preload("Products").First(&response).Error
	return response, err
}

func (s *storeServiceImpl) Report(storeId string, startDate time.Time, endDate time.Time) (web.StoreReport, error) {
	var report web.StoreReport
	err := database.DB.WithContext(s.ctx).Raw("select sum(t.total_product_price) as total_sales, sum(td.amount) as total_product_sold from transactions t join transaction_details td on td.transaction_id =t.id where t.store_id = ? and t.created_at between ? and ?", storeId, startDate, endDate).Scan(&report).Error
	return report, err
}
