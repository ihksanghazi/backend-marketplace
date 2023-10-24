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
