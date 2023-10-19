package services

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/ihksanghazi/backend-marketplace/model/web"
)

type TransactionService interface{}

type transactionServiceImpl struct{}

func NewTransactionService() TransactionService {
	return &transactionServiceImpl{}
}

func (t *transactionServiceImpl) CekOngkir() error {
	payload := strings.NewReader("origin=501&destination=114&weight=1700&courier=pos")

	req, _ := http.NewRequest("POST", os.Getenv("URL_ONGKIR")+"/cost", payload)

	req.Header.Add("key", os.Getenv("API_KEY_ONGKIR"))
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// Parse JSON and return it as a JSON response
	var data web.OngkirWebResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}
	return nil
}
