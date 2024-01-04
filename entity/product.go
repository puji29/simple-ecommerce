package entity

import "time"

type Product struct {
	ID            string    `json:"id"`
	ProductName   string    `json:"productName"`
	Description   string    `json:"description"`
	Price         int       `json:"price"`
	StockQuantity int       `json:"stockQuantity"`
	CreatedAT     time.Time `json:"createdAt"`
	UpdatedAT     time.Time `json:"updatedAt"`
	CategoryId    string    `json:"categoryId"`
	ImageId       string    `json:"imageId"`
}
