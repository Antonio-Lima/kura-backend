package model

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
	UserId   uint     `json:"-"`
	User     User     `json:"user" gorm:"foreignKey:UserId"`
	Category Category `json:"-" gorm:"foreignKey:CategoryId"`
	TransactionInput
}
