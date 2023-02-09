package domain

import "time"

type WithdrawalsRequest struct {
	Amount      float64 `form:"amount" binding:"required"`
	ReferenceID string  `form:"reference_id" binding:"required"`
}

type WithdrawalsResponse struct {
	Withdrawal WithdrawalsData `json:"withdrawal"`
}

type WithdrawalsData struct {
	ID          string       `json:"id"`
	WithdrawnBy string       `json:"withdrawn_by"`
	Amount      float64      `json:"amount"`
	Status      WalletStatus `json:"status"`
	WithdrawnAt time.Time    `json:"withdrawn_at"`
	Type        string       `json:"type"`
	ReferenceID string       `json:"reference_id"`
}
