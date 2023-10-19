package controllers

import (
	"github.com/gin-gonic/gin"
)

type TransactionController interface {
	CekOngir(c *gin.Context)
}

type tranctionControllerImpl struct{}

func NewTransactionController() TransactionController {
	return &tranctionControllerImpl{}
}

func (t *tranctionControllerImpl) CekOngir(c *gin.Context) {

}
