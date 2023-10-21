package services

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/ihksanghazi/backend-marketplace/database"
	"github.com/ihksanghazi/backend-marketplace/model/web"
	"github.com/ihksanghazi/backend-marketplace/utils"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"gorm.io/gorm"
)

type TransactionService interface {
	CekOngkir(cartId string, userId string, expedition string) (web.ExpeditionWebResponse, error)
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

func (t *transactionServiceImpl) Checkout(cartId string) (*coreapi.ChargeResponse, error) {

	var s coreapi.Client
	s.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)

	chargeReq := &coreapi.ChargeReq{
		PaymentType:  coreapi.PaymentTypeBankTransfer,
		BankTransfer: &coreapi.BankTransferDetails{Bank: midtrans.BankBca},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "12345",
			GrossAmt: 200000,
		},
		CustomerDetails: &midtrans.CustomerDetails{
			FName: "Azhi",
			Email: "asasa",
		},
		// Items: &[]midtrans.ItemDetails{
		// 	{Name: "Ongkir",Price: 20000,Category: }
		// },
	}

	database.DB.Transaction(func(tx *gorm.DB) error {
		// var cart domain.Cart
		// if err := tx.Model(cart).WithContext(t.ctx).Where("id = ?", cartId).First().Error; err != nil {
		// 	return err
		// }
		return nil
	})

	coreApiRes, err := s.ChargeTransaction(chargeReq)

	return coreApiRes, err
}
