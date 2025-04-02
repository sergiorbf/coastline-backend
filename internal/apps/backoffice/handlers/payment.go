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

func GetPayment(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, PaymentMethods)
}

func GetPaymentByID(context *gin.Context) {
	id := context.Param("id")
	for _, paymentMethod := range PaymentMethods {
		if paymentMethod.ID == id {
			context.IndentedJSON(http.StatusOK, paymentMethod)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "payment method not found"})
}

func CreatePayment(context *gin.Context) {
	var newPaymentMethod payment.PaymentMethod

	if err := context.BindJSON(&newPaymentMethod); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	PaymentMethods = append(PaymentMethods, newPaymentMethod)
	context.IndentedJSON(http.StatusCreated, newPaymentMethod)
}
