package models

import "time"

type Subscription struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	ProductID int       `json:"product_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserEmail string    `json:"user_email"`
}

func NewSubscription(userID, productID int, startDate, endDate time.Time, status, userEmail string) *Subscription {
	currentTime := time.Now()
	return &Subscription{
		UserID:    userID,
		ProductID: productID,
		StartDate: startDate,
		EndDate:   endDate,
		Status:    status,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		UserEmail: userEmail,
	}
}
