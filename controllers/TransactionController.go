package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/backend-marketplace/model/web"
)

type TransactionController interface {
	CekOngir(c *gin.Context)
}

type tranctionControllerImpl struct{}

func NewTransactionController() TransactionController {
	return &tranctionControllerImpl{}
}

func (t *tranctionControllerImpl) CekOngir(c *gin.Context) {
	payload := strings.NewReader("origin=501&destination=114&weight=1700&courier=pos")

	req, _ := http.NewRequest("POST", os.Getenv("URL_ONGKIR")+"/cost", payload)

	req.Header.Add("key", os.Getenv("API_KEY_ONGKIR"))
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	// Parse JSON and return it as a JSON response
	var data web.OngkirWebResponse
	if err := json.Unmarshal(body, &data); err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, data.Rajaongkir)
}
