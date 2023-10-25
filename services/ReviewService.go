package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/ihksanghazi/backend-marketplace/database"
	"github.com/ihksanghazi/backend-marketplace/model/domain"
	"github.com/ihksanghazi/backend-marketplace/model/web"
	"gorm.io/gorm"
)

type ReviewService interface {
	Create(userId string, productId string, req web.CreateReviewRequest) error
	Get(productId string, page int, limit int) (result []web.GetReviewResponse, totalPage int, err error)
	Update(reviewId string, req web.UpdateReviewRequest) (web.UpdateReviewRequest, error)
	Delete(reviewId string) error
}

type reviewServiceImpl struct {
	ctx context.Context
}

func NewReviewService(ctx context.Context) ReviewService {
	return &reviewServiceImpl{
		ctx: ctx,
	}
}

func (r *reviewServiceImpl) Create(userId string, productId string, req web.CreateReviewRequest) error {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		userId_UUID, err := uuid.Parse(userId)
		if err != nil {
			return err
		}
		productId_UUID, err := uuid.Parse(productId)
		if err != nil {
			return err
		}
		var review domain.Review
		review.UserId = userId_UUID
		review.ProductId = productId_UUID
		review.Comment = req.Comment
		if err := tx.WithContext(r.ctx).Create(&review).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r *reviewServiceImpl) Get(productId string, page int, limit int) (result []web.GetReviewResponse, totalPage int, err error) {
	var review []web.GetReviewResponse
	var model domain.Review
	var totalData int64
	offset := (page - 1) * limit
	Err := database.DB.Model(model).WithContext(r.ctx).Where("product_id = ?", productId).Count(&totalData).Limit(limit).Offset(offset).Find(&review).Error
	TotalPage := (int(totalData) + limit - 1) / limit
	return review, TotalPage, Err
}

func (r *reviewServiceImpl) Update(reviewId string, req web.UpdateReviewRequest) (web.UpdateReviewRequest, error) {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		var model domain.Review
		if err := tx.Model(model).WithContext(r.ctx).Where("id = ?", reviewId).Update("comment", req.Comment).First(&req).Error; err != nil {
			return err
		}
		return nil
	})
	return req, err
}

func (r *reviewServiceImpl) Delete(reviewId string) error {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		var review domain.Review
		if err := tx.Model(review).WithContext(r.ctx).Where("id = ?", reviewId).First(&review).Error; err != nil {
			return err
		}
		if err := tx.Model(review).WithContext(r.ctx).Where("id = ?", reviewId).Delete(&review).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
