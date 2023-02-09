package domain

import "time"

type DepositType string

const (
	Deposits   DepositType = "deposits"
	Withdrawal DepositType = "withdrawal"
)

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
	Type        DepositType  `json:"type"`
	ReferenceID string       `form:"reference_id" binding:"required"`
}
