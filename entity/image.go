package entity

import "time"

type Images struct {
	ID        string    `json:"id"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
