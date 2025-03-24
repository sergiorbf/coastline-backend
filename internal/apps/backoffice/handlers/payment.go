package handlers

import (
	"net/http"
	"webserver/internal/apps/backoffice/modules/domain/payment"

	"github.com/gin-gonic/gin"
)

var PaymentMethods = []payment.PaymentMethod{
	{
		ID:             "1",
		CardHolder:     "Sergio Filho",
		CardNumber:     "1234 5678 1234 5678",
		ExpirationDate: "12/23",
		CVV:            "123",
		CardType:       payment.CreditCard,
	},
}

func GetPaymentMethods(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, PaymentMethods)
}

func CreatePaymentMethod(c *gin.Context) {
	var newPaymentMethod payment.PaymentMethod

	if err := c.BindJSON(&newPaymentMethod); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	PaymentMethods = append(PaymentMethods, newPaymentMethod)
	c.IndentedJSON(http.StatusCreated, newPaymentMethod)
}
