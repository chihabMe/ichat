package models

type Product struct {
	Base
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
