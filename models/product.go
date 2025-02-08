package models

import "time"

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewProduct(name, description, category string, price float64) *Product {
	currentTime := time.Now()
	return &Product{
		Name:        name,
		Description: description,
		Price:       price,
		Category:    category,
		CreatedAt:   currentTime,
		UpdatedAt:   currentTime,
	}
}
