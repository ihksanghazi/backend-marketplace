package services

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/ihksanghazi/backend-marketplace/database"
	"github.com/ihksanghazi/backend-marketplace/model/domain"
	"github.com/ihksanghazi/backend-marketplace/model/web"
	"gorm.io/gorm"
)

type ProductService interface {
	Create(userId string, req web.CreateProductRequest) error
	Update(productId string, req web.UpdateProductRequest) (web.UpdateProductRequest, error)
	Delete(productId string) error
	Find(search string, page int, limit int) (result []web.FindProductResponse, totalPage int, err error)
	Get(productId string) (web.GetProductResponse, error)
}

type productServiceImpl struct {
	ctx context.Context
}

func NewProductService(ctx context.Context) ProductService {
	return &productServiceImpl{
		ctx: ctx,
	}
}

func (p *productServiceImpl) Create(userId string, req web.CreateProductRequest) error {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		// get store id
		var storeId string
		if err := tx.Raw("select s.id from stores s join users u on s.user_id =u.id where u.id = ?", userId).Scan(&storeId).Error; err != nil {
			return err
		}

		if storeId == "" {
			return errors.New("you don't have a shop")
		}

		storeId1, err := uuid.Parse(storeId)
		if err != nil {
			return err
		}

		var product domain.Product
		product.StoreId = storeId1
		product.ProductName = req.ProductName
		product.Description = req.Description
		product.Category = req.Category
		product.Stock = req.Stock
		product.Price = req.Price
		product.ImageUrl = req.ImageUrl
		product.WeightGram = req.WeightGram

		if err := tx.Model(product).WithContext(p.ctx).Create(&product).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (p *productServiceImpl) Update(productId string, req web.UpdateProductRequest) (web.UpdateProductRequest, error) {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		var product domain.Product
		product.ProductName = req.ProductName
		product.Description = req.Description
		product.Category = req.Category
		product.Stock = req.Stock
		product.Price = req.Price
		product.ImageUrl = req.ImageUrl
		product.WeightGram = req.WeightGram
		if err := tx.Model(product).WithContext(p.ctx).Where("id = ?", productId).Updates(product).First(&req).Error; err != nil {
			return err
		}
		return nil
	})

	return req, err
}

func (p *productServiceImpl) Delete(productId string) error {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		// cek product
		var product domain.Product
		if err := tx.Model(product).WithContext(p.ctx).Where("id = ?", productId).First(&product).Error; err != nil {
			return err
		}
		// delete product
		if err := tx.Model(product).WithContext(p.ctx).Where("id = ?", productId).Delete(&product).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (p *productServiceImpl) Find(search string, page int, limit int) (result []web.FindProductResponse, totalPage int, err error) {
	var product domain.Product
	var totalData int64
	var response []web.FindProductResponse
	offset := (page - 1) * limit
	Err := database.DB.Model(product).WithContext(p.ctx).Where("product_name ILIKE ? OR category ILIKE ? OR store_id = ?", "%"+search+"%", "%"+search+"%", search).Count(&totalData).Limit(limit).Offset(offset).Find(&response).Error
	TotalPage := (int(totalData) + limit - 1) / limit
	return response, TotalPage, Err
}

func (p *productServiceImpl) Get(productId string) (web.GetProductResponse, error) {
	var product domain.Product
	var response web.GetProductResponse
	err := database.DB.Model(product).WithContext(p.ctx).Where("id = ?", productId).Preload("Store").First(&response).Error
	return response, err
}
