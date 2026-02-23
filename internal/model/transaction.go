package model

import "github.com/google/uuid"

type TransactionType string

const (
	Income  TransactionType = "income"
	Expense TransactionType = "expense"
)

type TransactionInput struct {
	Amount      int64           `json:"amount" binding:"required,gt=0"`
	Description string          `json:"description" binding:"required,max=100"`
	CategoryId  uint            `json:"category_id" binding:"required"`
	Type        TransactionType `json:"type" binding:"required,oneof=income expense"`
}

type Transaction struct {
	Base
	UserId   uuid.UUID `json:"-"`
	User     User      `json:"user" gorm:"foreign_key:UserId"`
	Category Category  `json:"-" gorm:"foreign_key:CategoryId"`
	TransactionInput
}
