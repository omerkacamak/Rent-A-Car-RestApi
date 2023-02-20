package entity

import "time"

type Invoice struct {
	ID            int       `json:"id"`
	CreatedAt     time.Time ` json:"created_at"`
	UpdatedAt     time.Time ` json:"updated_at"`
	RentStartDate time.Time `json:"rentStartDate"`
	RentEndDate   time.Time `json:"rentEndDate"`
	TotalRentDays int       `json:"totalRentDays"`
	TotalPrice    float64   `json:"totalPrice"`

	CustomerID int       `json:"-"`
	Customer   *Customer `json:"customer,omitempty" gorm:"foreignKey:CustomerID"`

	PaymentID int      `json:"-"`
	Payment   *Payment `json:"payment,omitempty" gorm:"foreignKey:PaymentID"`
}
