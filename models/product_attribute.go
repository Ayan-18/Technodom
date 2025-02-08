package models

type ProductAttribute struct {
	ID        int    `json:"id"`
	ProductID int    `json:"product_id"`
	Name      string `json:"name"`
	Value     string `json:"value"`
}

func NewProductAttribute(productID int, name, value string) *ProductAttribute {
	return &ProductAttribute{
		ProductID: productID,
		Name:      name,
		Value:     value,
	}
}
