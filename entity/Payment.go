package entity

type Payment struct {
	ID         int     `json:"id"`
	TotalPrice float64 `json:"totalPrice"`

	CreditCardID int        `json:"creditcardid"`
	CreditCard   CreditCard `json:"-"`

	OrderID *int   `json:"orderid"`
	Order   *Order `json:"order,omitempty" gorm:"foreignKey:OrderID"`

	//omitemtpy eğer order boşsa json olarak döndürze doluysa döndürür
}
