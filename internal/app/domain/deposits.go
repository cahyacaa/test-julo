package domain

import "time"

type DepositsRequest struct {
	Amount      float64 `form:"amount" binding:"required"`
	ReferenceID string  `form:"reference_id" binding:"required"`
}

type DepositsResponse struct {
	Deposits DepositsData `json:"deposits"`
}

type DepositsData struct {
	ID          string       `json:"id"`
	DepositedBy string       `json:"deposited_by"`
	Amount      float64      `json:"amount"`
	Status      WalletStatus `json:"status"`
	DepositedAt time.Time    `json:"deposited_at"`
	Type        string       `json:"type"`
	ReferenceID string       `json:"reference_id"`
}
