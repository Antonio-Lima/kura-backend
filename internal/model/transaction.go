package model

import "time"

type TransactionType string

const (
	Income  TransactionType = "income"
	Expense TransactionType = "expense"
)

type TransactionInput struct {
	Amount      int64           `json:"amount" binding:"required,gt=0"`
	Description string          `json:"description" binding:"required,max=100"`
	Category    string          `json:"category"`
	Type        TransactionType `json:"type" binding:"required,oneof=income expense"`
}

type Transaction struct {
	ID        uint      `json:"id"`
	UserId    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	TransactionInput
}
