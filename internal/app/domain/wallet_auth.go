package domain

type WalletAuth struct {
	Token      string `json:"token"`
	IsDisabled bool   `json:"is_disabled"`
}
