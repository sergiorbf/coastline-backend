package handlers

import (
	"net/http"
	experience "webserver/internal/apps/backoffice/modules/domain/experiences"

	"github.com/gin-gonic/gin"
)

var experiences = []experience.Experience{
	{
		ID:          "1",
		Title:       "Surfing Lesson",
		Description: "Learn to surf with local instructors on the best beaches.",
		Price:       250,
		URL:         "/assets/experiences/surf-lesson.jpg",
	},
	{
		ID:          "2",
		Title:       "Bondinho Ride",
		Description: "Take a scenic ride on the famous Bondinho to Laranjeiras Beach, enjoying breathtaking views of the coastline.",
		Price:       100,
		URL:         "/assets/experiences/bondinho-laranjeiras.jpg",
	},
	{
		ID:          "3",
		Title:       "Pirate Ship Tour",
		Description: "Set sail on a thrilling pirate ship tour in Bombinhas, enjoying the coastal views and fun activities.",
		Price:       250,
		URL:         "/assets/experiences/pirate-ship-tour.jpg",
	},
}

func GetExperiences(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, experiences)
}
