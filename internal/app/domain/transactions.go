package domain

import "time"

type TransactionType string

const (
	Withdrawal TransactionType = "withdrawal"
	Deposits   TransactionType = "deposits"
)

type TransactionData struct {
	ID           string          `json:"id"`
	Status       string          `json:"status"`
	TransactedAt time.Time       `json:"transacted_at"`
	Type         TransactionType `json:"type"`
	Amount       float64         `json:"amount"`
	ReferenceID  string          `json:"reference_id"`
}

type TransactionsResponse struct {
	Transactions []TransactionData `json:"transactions"`
}
