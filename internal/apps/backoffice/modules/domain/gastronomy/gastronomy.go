package gastronomy

type Gastronomy struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Location    string `json:"location"`
	URL         string `json:"url"`
	Description string `json:"description"`
}
