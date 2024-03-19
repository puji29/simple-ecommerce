package entity

import "time"

type OrderDetail struct {
	ID          string    `json:"id"`
	OrderId     string    `json:"orderId"`
	ProductId   string    `json:"productId"`
	Quantity    int       `json:"quantity"`
	TotalAmount int       `json:"totalAmpunt"`
	CreatedAT   time.Time `json:"createdAt"`
	UpdatedAT   time.Time `json:"updatedAt"`
}
