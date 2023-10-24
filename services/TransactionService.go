package services

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ihksanghazi/backend-marketplace/database"
	"github.com/ihksanghazi/backend-marketplace/model/domain"
	"github.com/ihksanghazi/backend-marketplace/model/web"
	"github.com/ihksanghazi/backend-marketplace/utils"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"gorm.io/gorm"
)

type TransactionService interface {
	CekOngkir(cartId string, userId string, expedition string) (web.ExpeditionWebResponse, error)
	Checkout(cartId string, req web.CheckoutRequest, payment string) (*coreapi.ChargeResponse, error)
	Callback(orderId string) error
	GetByUserId(userId string) ([]web.GetTransactionResponse, error)
	GetByStoreId(storeId string) ([]web.GetTransactionResponse, error)
}

type transactionServiceImpl struct {
	ctx context.Context
}

func NewTransactionService(ctx context.Context) TransactionService {
	return &transactionServiceImpl{
		ctx: ctx,
	}
}

type cekOngkir struct {
	CityId     string
	Total_Gram string
}

var c coreapi.Client

func (t *transactionServiceImpl) CekOngkir(cartId string, userId string, expedition string) (web.ExpeditionWebResponse, error) {
	var result web.ExpeditionWebResponse

	// get origin city
	var cekOngkir cekOngkir
	if err := database.DB.WithContext(t.ctx).Raw("select s.city_id,c.total_gram from carts c join stores s on c.store_id = s.id where c.id = ? ", cartId).Scan(&cekOngkir).Error; err != nil {
		return result, err
	}
	// get destination city
	var destinationCity string
	if err := database.DB.WithContext(t.ctx).Raw("select u.city_id from users u where u.id = ? ", userId).Scan(&destinationCity).Error; err != nil {
		return result, err
	}

	payload := strings.NewReader("origin=" + cekOngkir.CityId + "&destination=" + destinationCity + "&weight=" + cekOngkir.Total_Gram + "&courier=" + expedition)

	req, err := http.NewRequest("POST", os.Getenv("URL_ONGKIR")+"/cost", payload)
	if err != nil {
		return result, err
	}

	req.Header.Add("key", os.Getenv("API_KEY_ONGKIR"))
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	body, err := utils.ResponseAPI(req)
	if err != nil {
		return result, err
	}

	// Parse JSON and return it as a JSON response
	var data web.Ongkir
	if err := json.Unmarshal(body, &data); err != nil {
		return result, err
	}

	var services []web.ServiceExpedition
	for _, items := range data.Rajaongkir.Results[0].Costs {
		service := web.ServiceExpedition{
			Service:     items.Service,
			Description: items.Description,
			Value:       items.Cost[0].Value,
			Etd:         items.Cost[0].Etd,
			Note:        items.Cost[0].Note,
		}
		services = append(services, service)

	}

	result.OriginDetails = data.Rajaongkir.OriginDetails
	result.DestinationDetails = data.Rajaongkir.DestinationDetails
	result.Weight = cekOngkir.Total_Gram
	result.Services = services

	return result, err
}

func (t *transactionServiceImpl) Checkout(cartId string, req web.CheckoutRequest, payment string) (*coreapi.ChargeResponse, error) {
	c.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)
	var chargeResponse *coreapi.ChargeResponse

	// transaction
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		var cart domain.Cart
		var responseCart web.GetCartResponse
		// get cart
		if err := tx.Model(cart).WithContext(t.ctx).Where("id = ?", cartId).Preload("Store").Preload("Items.Product").First(&responseCart).Error; err != nil {
			return err
		}
		// get user
		var user domain.User
		if err := tx.Model(user).WithContext(t.ctx).Where("id = ?", responseCart.UserId).First(&user).Error; err != nil {
			return err
		}

		totalProductPrice, err := strconv.Atoi(responseCart.Total)
		if err != nil {
			return err
		}
		// jumlah biaya seluruh product dan ongkos kirim
		totalPrice := totalProductPrice + req.Price
		orderId := "TRX-" + strings.Split(responseCart.Id.String(), "-")[0] + "-" + strconv.FormatInt(time.Now().Unix(), 10)
		chargeReq := &coreapi.ChargeReq{
			PaymentType:  coreapi.PaymentTypeBankTransfer,
			BankTransfer: &coreapi.BankTransferDetails{Bank: midtrans.Bank(payment)},
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  orderId,
				GrossAmt: int64(totalPrice),
			},
			CustomerDetails: &midtrans.CustomerDetails{
				FName: user.Username,
				Email: user.Email,
			},
		}
		coreApiRes, errCharge := c.ChargeTransaction(chargeReq)
		if errCharge != nil {
			return errCharge
		}

		chargeResponse = coreApiRes

		// insert to database
		var transaction domain.Transaction
		transaction.OrderId = orderId
		transaction.StoreId = responseCart.StoreId
		transaction.UserId = responseCart.UserId
		transaction.TransactionStatus = "pending"
		transaction.TotalProductPrice = responseCart.Total
		transaction.TotalPrice = strconv.Itoa(totalPrice)
		if err := tx.Model(transaction).WithContext(t.ctx).Create(&transaction).Error; err != nil {
			return err
		}
		var expedition domain.Expedition
		expedition.TransactionId = transaction.Id
		expedition.OriginCity = req.OriginCity
		expedition.DestinationCity = req.DestionationCity
		expedition.Courier = req.Courier
		expedition.WeightOnGram = req.WeightOnGram
		expedition.Service = req.Service
		expedition.Description = req.Description
		expedition.Price = req.Price
		if err := tx.Model(expedition).WithContext(t.ctx).Create(&expedition).Error; err != nil {
			return err
		}

		for _, item := range responseCart.Items {
			var transactionDetail domain.TransactionDetail
			transactionDetail.TransactionId = transaction.Id
			transactionDetail.ProductId = item.ProductId
			transactionDetail.Amount = item.Amount
			if err := tx.Model(transactionDetail).WithContext(t.ctx).Create(&transactionDetail).Error; err != nil {
				return err
			}
		}

		return nil
	})

	return chargeResponse, err
}

func (t *transactionServiceImpl) Callback(orderId string) error {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		var transaction domain.Transaction
		transactionStatusResp, err := c.CheckTransaction(orderId)
		if err != nil {
			return err
		} else {
			if transactionStatusResp != nil {
				// 5. Do set transaction status based on response from check transaction status
				if transactionStatusResp.TransactionStatus == "capture" {
					if transactionStatusResp.FraudStatus == "challenge" {
						// TODO set transaction status on your database to 'challenge'
						if err := tx.Model(transaction).WithContext(t.ctx).Where("order_id = ?", orderId).Update("transaction_status", "challenge").Error; err != nil {
							return err
						}
						// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
					} else if transactionStatusResp.FraudStatus == "accept" {
						// TODO set transaction status on your database to 'success'
						if err := tx.Model(transaction).WithContext(t.ctx).Where("order_id = ?", orderId).Update("transaction_status", "success").Error; err != nil {
							return err
						}
					}
				} else if transactionStatusResp.TransactionStatus == "settlement" {
					// TODO set transaction status on your databaase to 'success'
					if err := tx.Model(transaction).WithContext(t.ctx).Where("order_id = ?", orderId).Update("transaction_status", "success").Error; err != nil {
						return err
					}
				} else if transactionStatusResp.TransactionStatus == "deny" {
					// TODO you can ignore 'deny', because most of the time it allows payment retries
					if err := tx.Model(transaction).WithContext(t.ctx).Where("order_id = ?", orderId).Update("transaction_status", "deny").Error; err != nil {
						return err
					}
					// and later can become success
				} else if transactionStatusResp.TransactionStatus == "cancel" || transactionStatusResp.TransactionStatus == "expire" {
					// TODO set transaction status on your databaase to 'failure'
					if err := tx.Model(transaction).WithContext(t.ctx).Where("order_id = ?", orderId).Update("transaction_status", "failure").Error; err != nil {
						return err
					}
				} else if transactionStatusResp.TransactionStatus == "pending" {
					// TODO set transaction status on your databaase to 'pending' / waiting payment
					if err := tx.Model(transaction).WithContext(t.ctx).Where("order_id = ?", orderId).Update("transaction_status", "pending").Error; err != nil {
						return err
					}
				}
			}
		}
		return nil
	})
	return err
}

func (t *transactionServiceImpl) GetByUserId(userId string) ([]web.GetTransactionResponse, error) {
	var transaction []web.GetTransactionResponse
	err := database.DB.WithContext(t.ctx).Where("user_id = ?", userId).Preload("User.Region").Preload("Store.Region").Preload("Items.Product").Preload("Expedition").Find(&transaction).Error
	return transaction, err
}

func (t *transactionServiceImpl) GetByStoreId(storeId string) ([]web.GetTransactionResponse, error) {
	var transaction []web.GetTransactionResponse
	err := database.DB.WithContext(t.ctx).Where("store_id = ?", storeId).Preload("User.Region").Preload("Store.Region").Preload("Items.Product").Preload("Expedition").Find(&transaction).Error
	return transaction, err
}
