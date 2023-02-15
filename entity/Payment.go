package entity

type Payment struct {
	ID         int     `json:"id"`
	TotalPrice float64 `json:"totalPrice"`

	CreditCardID int        `json:"-"`
	CreditCard   CreditCard `json:"creditCard"`

	OrderID int   `json:"-"`
	Order   Order `json:"order" gorm:"foreignKey:OrderID"`
}
