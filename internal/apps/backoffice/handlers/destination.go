package handlers

import (
	"net/http"
	"webserver/internal/apps/backoffice/modules/domain/destination"

	"github.com/gin-gonic/gin"
)

var destinations = []destination.Destination{
	{
		ID:          "1",
		Name:        "Balneário Camboriú",
		URL:         "/assets/destinations/bc.jpg",
		Description: "Known as the Brazilian Dubai, famous for its nightlife and stunning beaches.",
	},
	{
		ID:          "2",
		Name:        "Florianópolis",
		URL:         "/assets/destinations/floripa-lagoinha.jpg",
		Description: "An island paradise with over 40 beaches, perfect for surfing and relaxation.",
	},
	{
		ID:          "3",
		Name:        "Itajaí",
		URL:         "/assets/destinations/brava.jpg",
		Description: "A charming coastal town known for its port, gastronomy, and nearby beaches.",
	},
	{
		ID:          "4",
		Name:        "Bombinhas",
		URL:         "/assets/destinations/bombinhas.jpg",
		Description: "A small paradise with crystal-clear waters, ideal for snorkeling and diving.",
	},
	{
		ID:          "5",
		Name:        "Itapema",
		URL:         "/assets/destinations/itapema.jpg",
		Description: "A cozy coastal city with golden sands and calm waters, perfect for families.",
	},
	{
		ID:          "6",
		Name:        "Praia do Rosa",
		URL:         "/assets/destinations/rosa.jpg",
		Description: "A stunning beach with lush landscapes, popular among surfers and nature lovers.",
	},
}

func GetDestinations(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, destinations)
}
