package entity

import "time"

type Category struct {
	ID           string    `json:"id"`
	CategoryName string    `json:"categoryName"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
