package domain

type InitWalletAccountRequest struct {
	CustomerXID string `form:"customer_xid" binding:"required"`
}
type InitWalletAccountResponse struct {
	Token string `json:"token" `
}
