package services

import (
	"context"
	"errors"
	"strconv"

	"github.com/google/uuid"
	"github.com/ihksanghazi/backend-marketplace/database"
	"github.com/ihksanghazi/backend-marketplace/model/domain"
	"github.com/ihksanghazi/backend-marketplace/model/web"
	"gorm.io/gorm"
)

type CartService interface {
	Add(productId string, amount string, userId string) error
	Get(userId string) ([]web.GetCartResponse, error)
	DeleteCart(cartId string) error
	UpdateCartItem(itemId string, qty int) error
	DeleteCartItem(itemId string) error
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

		// cek apakah user bukan penjual dari barang ini
		var productUserId string
		if err := tx.WithContext(c.ctx).Raw("select u.id from users u join stores s on u.id = s.user_id where s.id = ?", product.StoreId).Scan(&productUserId).Error; err != nil {
			return err
		}
		if productUserId == userId {
			return errors.New("you are the selller of this item")
		}

		// cek apakah product masih tersedia
		productStock, err := strconv.Atoi(product.Stock)
		if err != nil {
			return err
		}
		productAmount, err := strconv.Atoi(amount)
		if err != nil {
			return err
		}
		if productStock-productAmount < 0 {
			return errors.New("product is not available")
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

func (c *cartServiceImpl) DeleteCart(cartId string) error {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		var cart domain.Cart
		if err := tx.Model(cart).WithContext(c.ctx).Where("id = ?", cartId).First(&cart).Error; err != nil {
			return err
		}
		var cartDetail domain.CartDetail
		if err := tx.Model(cartDetail).WithContext(c.ctx).Where("cart_id = ?", cartId).Delete(&cartDetail).Error; err != nil {
			return err
		}
		if err := tx.Model(cart).WithContext(c.ctx).Where("id = ?", cartId).Delete(&cart).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (c *cartServiceImpl) UpdateCartItem(itemId string, qty int) error {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		var cartDetail domain.CartDetail
		if err := tx.Model(cartDetail).WithContext(c.ctx).Where("id = ?", itemId).First(&cartDetail).Error; err != nil {
			return err
		}
		if err := tx.Model(cartDetail).WithContext(c.ctx).Where("id = ?", itemId).Update("amount", qty).Error; err != nil {
			return err
		}
		// hitung ulang total cart
		var total int
		if err := tx.WithContext(c.ctx).Raw("select sum(cd.amount*p.price) from cart_details cd join products p ON cd.product_id =p.id where cd.cart_id = ?", cartDetail.CartId).Scan(&total).Error; err != nil {
			return err
		}
		var cart domain.Cart
		if err := tx.Model(cart).WithContext(c.ctx).Where("id = ?", cartDetail.CartId).Update("total", total).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (c *cartServiceImpl) DeleteCartItem(itemId string) error {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		var cartDetail domain.CartDetail
		var cart domain.Cart
		if err := tx.Model(cartDetail).WithContext(c.ctx).Where("id = ?", itemId).First(&cartDetail).Error; err != nil {
			return err
		}
		cart.Id = cartDetail.CartId
		if err := tx.Model(cartDetail).WithContext(c.ctx).Where("id = ?", itemId).Delete(&cartDetail).Error; err != nil {
			return err
		}

		// jika cart detail sudah kosong maka hapus cartnya jika ada maka hitung ulang cartnya
		cartDetail.Id = uuid.Nil
		if err := tx.Model(cartDetail).WithContext(c.ctx).Where("cart_id = ?", cartDetail.CartId).Find(&cartDetail).Error; err != nil {
			return err
		}

		if cartDetail.Id != uuid.Nil {
			// hitung ulang total cart
			var total int
			if err := tx.WithContext(c.ctx).Raw("select sum(cd.amount*p.price) from cart_details cd join products p ON cd.product_id =p.id where cd.cart_id = ?", cartDetail.CartId).Scan(&total).Error; err != nil {
				return err
			}

			if err := tx.Model(cart).WithContext(c.ctx).Where("id = ?", cartDetail.CartId).Update("total", total).Error; err != nil {
				return err
			}
		} else {
			if err := tx.Model(cart).WithContext(c.ctx).Where("id = ?", cart.Id).Delete(&cart).Error; err != nil {
				return err
			}
		}
		return nil
	})
	return err
}
