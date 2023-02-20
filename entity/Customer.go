package entity

type Customer struct {
	ID             int    `json:"id" gorm:"primaryKey"`
	FirstName      string `json:"firstname" binding:"required" validate:"lent" `
	LastName       string `json:"lastname" binding:"required" validate:"deneme"`
	AddressID      int    `json:"-"`
	Address        string `json:"address"`
	IdentityNumber string `json:"identitynumber"`

	Orders []Order `json:",omitempty" gorm:"foreignKey:CustomerID"`
}
