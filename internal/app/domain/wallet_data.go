package domain

import "time"

type WalletStatus string

const (
	Enabled  WalletStatus = "enabled"
	Disabled WalletStatus = "disabled"
)

type WalletResponse struct {
	Wallet struct {
		ID        string       `json:"id"`
		OwnedBy   string       `json:"owned_by"`
		Balance   float64      `json:"balance"`
		Status    WalletStatus `json:"status"`
		EnabledAt time.Time    `json:"enabled_at"`
	} `json:"wallet"`
}

type WalletData struct {
	Balance    float64   `json:"balance"`
	CustomerID string    `json:"customer_id"`
	IsDisabled bool      `json:"is_disabled"`
	EnabledAt  time.Time `json:"enabled_at"`
}
