package model

import "time"

type CategoryInput struct {
	Category string `json:"category" binding:"required"`
	Color    string `json:"color" binding:"required,hexcolor"`
}

type Category struct {
	ID        uint      `json:"id"`
	UserId    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CategoryInput
}
