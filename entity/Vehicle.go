package entity

import "github.com/omerkacamak/rentacar-golang/enum"

type Vehicle struct {
	ID         int           `json:"id" gorm:"primaryKey"`
	Model      string        `json:"model"`
	CarYear    string        `json:"caryear"`
	Category   enum.Category `json:"category"`
	Bluetooth  bool          `json:"bluetooth"`
	Navigation bool          `json:"navigation"`

	BrandID      int          `json:"brandid"`
	VehicleBrand VehicleBrand `json:"brand" gorm:"foreignKey:BrandID"`

	// ColorID VehicleColor `json:"-"`
	// Color   VehicleColor `json:"color" gorm:"foreignKey:ColorID;references:ID"`
}
