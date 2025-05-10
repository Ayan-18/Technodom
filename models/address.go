package models

type Address struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Street    string `json:"street"`
	City      string `json:"city"`
	State     string `json:"state"`
	ZipCode   string `json:"zip_code"`
	Country   string `json:"country"`
	IsPrimary bool   `json:"is_primary"` // Основной адрес или нет
}

func NewAddress(userID int, street, city, state, zipCode, country string, isPrimary bool) *Address {
	return &Address{
		UserID:    userID,
		Street:    street,
		City:      city,
		State:     state,
		ZipCode:   zipCode,
		Country:   country,
		IsPrimary: isPrimary,
	}
}
