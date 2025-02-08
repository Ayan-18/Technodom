package models

import "time"

type Review struct {
	ID        int       `json:"id"`
	ProductID int       `json:"product_id"`
	UserID    int       `json:"user_id"`
	Rating    int       `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}

func NewReview(productID, userID, rating int, comment string) *Review {
	return &Review{
		ProductID: productID,
		UserID:    userID,
		Rating:    rating,
		Comment:   comment,
		CreatedAt: time.Now(),
	}
}
