package database

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/ihksanghazi/backend-marketplace/model/domain"
)

type responseProvince struct {
	Rajaongkir rajaongkir `json:"rajaongkir"`
}

type rajaongkir struct {
	Query   interface{} `json:"query"`
	Status  interface{} `json:"status"`
	Results []city      `json:"results"`
}

type city struct {
	CityId     string `json:"city_id"`
	ProvinceId string `json:"province_id"`
	Province   string `json:"province"`
	Type       string `json:"type"`
	CityName   string `json:"city_name"`
	PostalCode string `json:"postal_code"`
}

func Seeder() {
	req, err := http.NewRequest("GET", os.Getenv("URL_ONGKIR")+"/city", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("key", os.Getenv("API_KEY_ONGKIR"))
	// req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

	// Parse JSON and return it as a JSON response
	var data responseProvince
	json.Unmarshal(body, &data)

	for _, item := range data.Rajaongkir.Results {
		var city domain.City
		city.Id = item.CityId
		city.CityName = item.CityName
		city.PostalCode = item.PostalCode
		city.ProvinceId = item.ProvinceId
		city.Type = item.Type
		if err := DB.Model(city).Create(&city).Error; err != nil {
			panic(err)
		}
	}

}
