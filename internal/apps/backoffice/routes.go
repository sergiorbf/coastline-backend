package backoffice

import (
	"webserver/internal/apps/backoffice/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() {
	router := gin.Default()

	// Destination
	router.GET("/destination", handlers.GetDestinations)

	// Gastronomy
	router.GET("/gastronomy", handlers.GetGastronomy)

	// Experiences
	router.GET("/experiences", handlers.GetExperiences)

	// Profile
	router.GET("/account-settings/profile", handlers.GetProfile)

	// Payment
	paymentGroup := router.Group("/payments")
	paymentGroup.GET("", handlers.GetPayment)
	paymentGroup.GET("/:id", handlers.GetPaymentByID)
	paymentGroup.POST("/create", handlers.CreatePayment)
	paymentGroup.PUT("/update/:id", handlers.UpdatePayment)
	paymentGroup.DELETE("/delete/:id", handlers.DeletePayment)

	router.Run("localhost:8080")
}
