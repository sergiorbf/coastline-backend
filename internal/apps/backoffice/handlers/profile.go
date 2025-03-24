package handlers

import (
	"net/http"
	"webserver/internal/apps/backoffice/modules/domain/profile"

	"github.com/gin-gonic/gin"
)

var Profile = profile.Profile{
	ID:         "1",
	Name:       "Sergio",
	LastName:   "Filho",
	Username:   "srbf",
	Email:      "srbf@mail.com",
	Phone:      "123456789",
	MonthBirth: 3,
	DayBirth:   9,
	YearBirth:  2002,
	ProfilePic: nil,
	Bio:        "I'm a software developer",
	Country:    1,
	Address:    "Rua 1",
	Address2:   "Apto 1",
	ZipCode:    "123456",
	City:       1,
	State:      1,
}

func GetProfile(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Profile)
}
