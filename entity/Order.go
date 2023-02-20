package entity

import "time"

type Order struct {
	ID          int `json:"id"`
	KmStart     float64
	KmEnd       float64
	StartDate   time.Time `json:"startdate"`
	EndDate     time.Time `json:"enddate"`
	OrderStatus bool      `json:"orderstatus"`
	OrderPrice  float64   `json:"orderprice"`

	RentCityID int   `json:"-"`
	RentCity   *City `json:"rentCity,omitempty" gorm:"foreignKey:RentCityID"`

	ReturnCityID int   `json:"-"`
	ReturnCity   *City `json:"returnCity,omitempty" gorm:"foreignKey:ReturnCityID"`

	VehicleID int      `json:"-"`
	Vehicle   *Vehicle `json:"vehicle,omitempty" gorm:"foreignKey:VehicleID"`

	CustomerID int       `json:"CustomerID"`
	Customer   *Customer `json:"customer,omitempty" gorm:"foreignKey:CustomerID"`

	CreatedAt time.Time ` json:"created_at"`
	UpdatedAt time.Time ` json:"updated_at"`
}
