package domain

import "time"

type WalletStatus string

const (
	Enabled  WalletStatus = "enabled"
	Disabled WalletStatus = "disabled"
)

type WalletResponse struct {
	Wallet WalletResponseData `json:"wallet"`
}

type WalletResponseData struct {
	ID         string       `json:"id"`
	OwnedBy    string       `json:"owned_by"`
	Balance    float64      `json:"balance"`
	Status     WalletStatus `json:"status"`
	EnabledAt  *time.Time   `json:"enabled_at,omitempty"`
	DisabledAt *time.Time   `json:"disabled_at,omitempty"`
}

type WalletData struct {
	ID         string    `json:"id"`
	Balance    float64   `json:"balance"`
	CustomerID string    `json:"customer_id"`
	IsDisabled bool      `json:"is_disabled"`
	EnabledAt  time.Time `json:"enabled_at"`
	DisabledAt time.Time `json:"disabled_at"`
}

type EnabledWallet struct {
	ID        string       `json:"id"`
	OwnedBy   string       `json:"owned_by"`
	Status    WalletStatus `json:"status"`
	EnabledAt time.Time    `json:"enabled_at"`
	Balance   float64      `json:"balance"`
}

type DisabledWallet struct {
	ID         string       `json:"id"`
	OwnedBy    string       `json:"owned_by"`
	Status     WalletStatus `json:"status"`
	DisabledAt time.Time    `json:"disabled_at"`
	Balance    float64      `json:"balance"`
}
