package entity

type Payment struct {
	ID         int     `json:"id"`
	TotalPrice float64 `json:"totalPrice"`

	CreditCardID int        `json:"creditcardid"`
	CreditCard   CreditCard `json:"credit,omitempty"`

	OrderID int   `json:"orderid"`
	Order   Order `json:"order,omitempty" gorm:"foreignKey:OrderID"`
}
