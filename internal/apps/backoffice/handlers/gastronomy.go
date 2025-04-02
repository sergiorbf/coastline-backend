package handlers

import (
	"net/http"
	"webserver/internal/apps/backoffice/modules/domain/gastronomy"

	"github.com/gin-gonic/gin"
)

var restaurants = []gastronomy.Gastronomy{
	{
		ID:          "1",
		Name:        "Restaurante O Barco",
		Location:    "Balneário Camboriú, SC",
		URL:         "/assets/gastronomy/barco.jpg",
		Description: "Known for its fresh seafood dishes and beachside location.",
	},
	{
		ID:          "2",
		Name:        "Ostradamus",
		Location:    "Florianópolis, SC",
		URL:         "/assets/gastronomy/ostradamus.jpg",
		Description: "A must-visit spot for oyster lovers, located in the charming Ribeirão da Ilha.",
	},
	{
		ID:          "3",
		Name:        "Le Poulet Rouge",
		Location:    "Balneário Camboriú, SC",
		URL:         "/assets/gastronomy/le-poulet-rouge.jpg",
		Description: "A charming French-inspired bistro offering delicious poultry dishes and a cozy atmosphere.",
	},
	{
		ID:          "4",
		Name:        "Quintal do Félix",
		Location:    "Navegantes, SC",
		URL:         "/assets/gastronomy/quintal-do-felix.jpg",
		Description: "A charming local restaurant in Navegantes offering fresh seafood and regional dishes in a cozy and rustic atmosphere.",
	},
	{
		ID:          "5",
		Name:        "Sushiarte Praia Brava",
		Location:    "Itajaí, SC",
		URL:         "/assets/gastronomy/sushiarte.jpg",
		Description: "A sushi restaurant offering fresh and creative Japanese dishes with an artistic touch, located in the stunning Praia Brava.",
	},
}

func GetGastronomy(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, restaurants)
}
