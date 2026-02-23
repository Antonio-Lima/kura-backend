package model

type UserBase struct {
	Email  string `json:"email" binding:"required,email" gorm:"unique"`
	Name   string `json:"name" binding:"required,min=3"`
	Avatar string `json:"avatar" binding:"omitempty,url"`
}

type UserInput struct {
	UserBase
	Password             string `json:"password" binding:"required,min=8"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required,eqfield=Password"`
}

type User struct {
	Base
	Password string `json:"-"`
	UserBase
}
