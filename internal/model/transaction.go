package model

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
	Base
	UserId uint `json:"user_id"`
	TransactionInput
}
