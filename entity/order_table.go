package entity

import "time"

type OrderTable struct {
	ID          string    `json:"id"`
	UserId      string    `json:"userId"`
	OrderDate   time.Time `json:"orderDate"`
	TotalAmount int       `json:"totalAmount"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
