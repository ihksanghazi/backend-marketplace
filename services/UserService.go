package services

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/ihksanghazi/backend-marketplace/database"
	"github.com/ihksanghazi/backend-marketplace/model/domain"
	"github.com/ihksanghazi/backend-marketplace/model/web"
	"github.com/ihksanghazi/backend-marketplace/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface{
	Register(req web.RegisterRequest) (web.RegisterResponse,error)
	Login(req web.LoginRequest) (refreshToken string, accessToken string, err error)
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

func (u *userServiceImpl) Login(req web.LoginRequest) (refreshToken string, accessToken string, err error) {

	var r_token,a_token string
	Err:=database.DB.Transaction(func(tx *gorm.DB) error {
		var user domain.User
		// find user
		if err := tx.Model(user).WithContext(u.ctx).Where("email = ?",req.Email).First(&user).Error;err!= nil {
			return err
		}

		// cek password
		if err:= bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(req.Password)); err!= nil {
			return errors.New("wrong password")
		}

		// generete Refresh Token
		refreshToken,err:= utils.GenerateToken(os.Getenv("REFRESH_TOKEN"),time.Now().Add(24 * time.Hour),user.Id.String(),user.Email,user.Username)
		if err!= nil {
			return err
		}
		r_token = refreshToken

		// generate Access Token
		accessToken,err:= utils.GenerateToken(os.Getenv("ACCESS_TOKEN"),time.Now().Add(30 * time.Second),user.Id.String(),user.Email,user.Username)
		if err!= nil {
			return err
		}
		a_token = accessToken

		// update user
		if err:= tx.Model(user).WithContext(u.ctx).Where("id = ?",user.Id).Update("refresh_token",refreshToken).Error; err!= nil {
			return err
		}
		return nil
	})

	return r_token,a_token,Err
}