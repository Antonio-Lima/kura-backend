package model

type CategoryInput struct {
	Category string `json:"category" binding:"required"`
	Color    string `json:"color" binding:"required,hexcolor"`
}

type Category struct {
	Base
	UserId uint `json:"user_id"`
	CategoryInput
}
