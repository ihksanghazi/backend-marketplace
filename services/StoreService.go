package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/ihksanghazi/backend-marketplace/database"
	"github.com/ihksanghazi/backend-marketplace/model/domain"
	"github.com/ihksanghazi/backend-marketplace/model/web"
	"gorm.io/gorm"
)

type StoreService interface{
	Create(userId string,req web.CreateStoreRequest) (web.CreateStoreRequest,error)
}

type storeServiceImpl struct {
	ctx context.Context
}

func NewStoreService(ctx context.Context) StoreService {
	return &storeServiceImpl{
		ctx:ctx,
	}
}

func (s *storeServiceImpl) Create(userId string,req web.CreateStoreRequest) (web.CreateStoreRequest,error) {
	err:=database.DB.Transaction(func(tx *gorm.DB) error {
		UserId,err:=uuid.Parse(userId)
		if err != nil {
			return err
		}
		
		var store domain.Store
		store.UserId = UserId
		store.StoreName = req.StoreName
		store.Description = req.Description
		store.Category = req.Category
		store.ImageUrl = req.ImageUrl
		
		if err:= tx.Model(store).WithContext(s.ctx).Create(&store).First(&req).Error; err!=nil {
			return err
		}
		return nil
	})
	return req,err
}