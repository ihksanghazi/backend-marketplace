package services

import (
	"context"

	"github.com/ihksanghazi/backend-marketplace/database"
	"github.com/ihksanghazi/backend-marketplace/model/domain"
	"github.com/ihksanghazi/backend-marketplace/model/web"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface{
	Register(req web.RegisterRequest) (web.RegisterResponse,error)
}

type userServiceImpl struct{
	ctx context.Context
}

func NewUserService(ctx context.Context) UserService {
	return &userServiceImpl{
		ctx:ctx,
	}
}

func (u *userServiceImpl) Register(req web.RegisterRequest) (web.RegisterResponse,error) {
	
	var res web.RegisterResponse
	
	err:= database.DB.Transaction(func(tx *gorm.DB) error {
		// hash password
		password,err:=bcrypt.GenerateFromPassword([]byte(req.Password),bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		var user domain.User
		user.Username = req.Username
		user.Email = req.Email
		user.Password = string(password)
		user.Address = req.Address
		user.ImageUrl = req.ImageUrl

		if err:= tx.Model(user).WithContext(u.ctx).Create(&user).First(&res).Error;err!= nil {
			return err
		}
		return nil
	})

	return res,err
}