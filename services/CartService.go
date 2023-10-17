package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/ihksanghazi/backend-marketplace/database"
	"github.com/ihksanghazi/backend-marketplace/model/domain"
	"github.com/ihksanghazi/backend-marketplace/model/web"
	"gorm.io/gorm"
)

type CartService interface {
	Add(productId string, amount string, userId string) error
	Get(userId string) ([]web.GetCartResponse, error)
}

type cartServiceImpl struct {
	ctx context.Context
}

func NewCartService(ctx context.Context) CartService {
	return &cartServiceImpl{
		ctx: ctx,
	}
}

func (c *cartServiceImpl) Add(productId string, amount string, userId string) error {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		// cari product berdasarkan id
		var product domain.Product
		if err := tx.Model(product).WithContext(c.ctx).Where("id = ?", productId).First(&product).Error; err != nil {
			return err
		}

		// cari cart berdasarkan user id dan store id
		var cart domain.Cart
		if err := tx.Model(cart).WithContext(c.ctx).Where("user_id = ? AND store_id = ?", userId, product.StoreId).Find(&cart).Error; err != nil {
			return err
		}
		// cek cart apakah sudah ada jika kosong maka buat cart baru
		if cart.Id == uuid.Nil {
			cart.UserId = uuid.MustParse(userId)
			cart.StoreId = product.StoreId
			cart.Total = "0"
			if err := tx.Model(cart).WithContext(c.ctx).Create(&cart).Error; err != nil {
				return err
			}
		}

		// cari cart detail berdasarkan cart id dan product id
		var cartDetail domain.CartDetail
		if err := tx.Model(cartDetail).WithContext(c.ctx).Where("cart_id = ? AND product_id = ?", cart.Id, product.Id).Find(&cartDetail).Error; err != nil {
			return err
		}
		// cek apakah cart detail sudah ada jika kosong maka buat cart detail baru
		if cartDetail.Id == uuid.Nil {
			cartDetail.CartId = cart.Id
			cartDetail.ProductId = product.Id
			cartDetail.Amount = amount
			if err := tx.Model(cartDetail).WithContext(c.ctx).Create(&cartDetail).Error; err != nil {
				return err
			}
			// jika ada maka update amountnya saja
		} else {
			if err := tx.Model(cartDetail).WithContext(c.ctx).Where("cart_id = ? AND product_id = ?", cart.Id, product.Id).Update("amount", gorm.Expr("amount + ?", amount)).Error; err != nil {
				return err
			}
		}

		// hitung ulang total cart
		var total int
		if err := tx.WithContext(c.ctx).Raw("select sum(cd.amount*p.price) from cart_details cd join products p ON cd.product_id =p.id where cd.cart_id = ?", cart.Id).Scan(&total).Error; err != nil {
			return err
		}
		if err := tx.Model(cart).WithContext(c.ctx).Where("id = ?", cart.Id).Update("total", total).Error; err != nil {
			return err
		}
		return nil
	})

	return err
}

func (c *cartServiceImpl) Get(userId string) ([]web.GetCartResponse, error) {
	var cart domain.Cart
	var response []web.GetCartResponse
	err := database.DB.Model(cart).WithContext(c.ctx).Where("user_id = ?", userId).Preload("Store").Preload("Products.Detail").Find(&response).Error
	return response, err
}
