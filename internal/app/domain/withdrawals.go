package domain

import "time"

type WithdrawalsType string

const (
	Withdrawal WithdrawalsType = "withdrawal"
)

type WithdrawalsRequest struct {
	Amount      float64 `form:"amount" binding:"required"`
	ReferenceID string  `form:"reference_id" binding:"required"`
}

type WithdrawalsResponse struct {
	Withdrawal WithdrawalsData `json:"withdrawal"`
}

type WithdrawalsData struct {
	ID          string          `json:"id"`
	WithdrawnBy string          `json:"withdrawn_by"`
	Amount      float64         `json:"amount"`
	Status      WalletStatus    `json:"status"`
	WithdrawnAt time.Time       `json:"withdrawn_at"`
	Type        WithdrawalsType `json:"type"`
	ReferenceID string          `form:"reference_id" binding:"required"`
}
