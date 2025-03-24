package profile

import "image"

type Profile struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	LastName   string      `json:"lastName"`
	Username   string      `json:"username"`
	Email      string      `json:"email"`
	Phone      string      `json:"phone"`
	MonthBirth int         `json:"monthBirth"`
	DayBirth   int         `json:"dayBirth"`
	YearBirth  int         `json:"yearBirth"`
	ProfilePic image.Image `json:"profilePic"`
	Bio        string      `json:"bio"`
	Country    int         `json:"country"`
	Address    string      `json:"address"`
	Address2   string      `json:"address2"`
	ZipCode    string      `json:"zipCode"`
	City       int         `json:"city"`
	State      int         `json:"state"`
}
