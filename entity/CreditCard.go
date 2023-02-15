package entity

type CreditCard struct {
	ID             int    `json:"id"`
	NickName       string `json:"name"`
	CardNumber     string `json:"cardNumber"`
	CardHolderName string `json:"cardHolderName"`
	ExpiryDate     string `json:"expiryDate"`
	SecurityCode   string `json:"securityCode"`
}
