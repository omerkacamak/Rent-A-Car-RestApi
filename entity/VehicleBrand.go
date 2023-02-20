package entity

type VehicleBrand struct {
	ID   int    `json:"id"`
	Name string `json:"name"`

	Vehicles []Vehicle `json:",omitempty" gorm:"foreignKey:BrandID"`
}
