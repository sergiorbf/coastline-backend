package payment

type CardType string

const (
	DebitCard  CardType = "debit"
	CreditCard CardType = "credit"
)

type PaymentMethod struct {
	ID             string   `json:"id"`
	CardHolder     string   `json:"cardHolder"`
	CardNumber     string   `json:"cardNumber"`
	ExpirationDate string   `json:"expirationDate"`
	CVV            string   `json:"cvv"`
	CardType       CardType `json:"cardType"`
}
