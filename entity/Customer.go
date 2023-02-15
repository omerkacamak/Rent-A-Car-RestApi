package entity

type Customer struct {
	ID             int    `json:"id"`
	FirstName      string `json:"firstname" validate:"required"`
	LastName       string `json:"lastname" validate:"required"`
	AddressID      int    `json:"-"`
	Address        string `json:"address"`
	IdentityNumber string `json:"identitynumber"`

	Orders []Order `gorm:"foreignKey:CustomerID"`
}
