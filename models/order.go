package models

import "time"

type Order struct {
	ID              int     `json:"id"`
	UserID          int     `json:"user_id"`
	Products        []int   `json:"products"`
	TotalPrice      float64 `json:"total_price"`
	Status          string  `json:"status"`
	PaymentMethod   string  `json:"payment_method"`
	ShippingAddress string  `json:"shipping_address"`
	//StatusLog       []StatusLog `json:"status_log"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewOrder(userID int, products []int, totalPrice float64, status, paymentMethod, shippingAddress string) *Order {
	currentTime := time.Now()
	order := &Order{
		UserID:          userID,
		Products:        products,
		TotalPrice:      totalPrice,
		Status:          status,
		PaymentMethod:   paymentMethod,
		ShippingAddress: shippingAddress,
		//StatusLog:       []StatusLog{},
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	//statusLog := NewStatusLog(order.ID, "order", "", "created", "system")

	//order.StatusLog = append(order.StatusLog, *statusLog)

	return order
}
