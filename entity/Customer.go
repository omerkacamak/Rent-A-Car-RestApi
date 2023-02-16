package entity

type Customer struct {
	ID             int    `json:"id" gorm:"primaryKey"`
	FirstName      string `json:"firstname" binding:"required" validate:"required,lent" `
	LastName       string `json:"lastname" binding:"required" validate:"required,deneme"`
	AddressID      int    `json:"-"`
	Address        string `json:"address"`
	IdentityNumber string `json:"identitynumber"`

	Orders []Order `gorm:"foreignKey:CustomerID"`
}
