package entity

import "time"

type User struct {
	ID           string    `json:"id"`
	Username     string    `json:"username"`
	UserRole     string    `json:"userRole"`
	Addres       string    `json:"address"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"passwordHash"`
	Created_at   time.Time `json:"createdAt"`
	Updated_at   time.Time `json:"updatedAt"`
}
