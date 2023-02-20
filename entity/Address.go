package entity

type Address struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	District    string `json:"district"`
	Description string `json:"description"`
}
