package services

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/ihksanghazi/backend-marketplace/database"
	"github.com/ihksanghazi/backend-marketplace/model/web"
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

func (t *transactionServiceImpl) CekOngkir(cartId string, userId string, expedition string) (web.ExpeditionWebResponse, error) {
	var result web.ExpeditionWebResponse

	// get origin city
	var originCity, destinationCity string
	if err := database.DB.WithContext(t.ctx).Raw("select s.city_id from carts c join stores s on c.store_id = s.id where c.id = ? ", cartId).Scan(&originCity).Error; err != nil {
		return result, err
	}
	// get destination city
	if err := database.DB.WithContext(t.ctx).Raw("select u.city_id from users u where u.id = ? ", userId).Scan(&destinationCity).Error; err != nil {
		return result, err
	}

	payload := strings.NewReader("origin=" + originCity + "&destination=" + destinationCity + "&weight=1000&courier=" + expedition)

	req, err := http.NewRequest("POST", os.Getenv("URL_ONGKIR")+"/cost", payload)
	if err != nil {
		return result, err
	}

	req.Header.Add("key", os.Getenv("API_KEY_ONGKIR"))
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
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
		//
		for _, cost := range items.Cost {
			service := web.ServiceExpedition{
				Service:     items.Service,
				Description: items.Description,
				Value:       cost.Value,
				Etd:         cost.Etd,
				Note:        cost.Note,
			}
			services = append(services, service)
		}
	}

	result.OriginDetails = data.Rajaongkir.OriginDetails
	result.DestinationDetails = data.Rajaongkir.DestinationDetails
	result.Services = services

	return result, err
}
