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
	router.GET("/payments", handlers.GetPaymentMethods)
	router.POST("/payments/create", handlers.CreatePaymentMethod)

	router.Run("localhost:8080")
}
