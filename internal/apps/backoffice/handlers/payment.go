package handlers

import (
	"net/http"
	"strconv"
	"webserver/internal/apps/backoffice/modules/domain/payment"

	"github.com/gin-gonic/gin"
)

var PaymentMethods = map[int]payment.PaymentMethod{
	1: {
		ID:             1,
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
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

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

	newPaymentMethod.ID = len(PaymentMethods) + 1
	PaymentMethods[len(PaymentMethods)+1] = newPaymentMethod

	context.IndentedJSON(http.StatusCreated, newPaymentMethod)
}

func UpdatePayment(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var updatedPaymentMethod payment.PaymentMethod

	if err := context.BindJSON(&updatedPaymentMethod); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, paymentMethod := range PaymentMethods {
		if paymentMethod.ID == id {
			PaymentMethods[i] = updatedPaymentMethod
			context.IndentedJSON(http.StatusOK, updatedPaymentMethod)
			return
		}
	}

	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "payment method not found"})

}

func DeletePayment(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	for i, paymentMethod := range PaymentMethods {
		if paymentMethod.ID == id {
			delete(PaymentMethods, i)
			context.IndentedJSON(http.StatusOK, gin.H{"message": "payment method deleted"})
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "payment method not found"})
}
