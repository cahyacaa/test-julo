package service

import "context"

type Wallet interface {
	InitWallet(ctx context.Context, customerID string) (token string, err error)
	EnableWallet(ctx context.Context, customerID string) (err error)
}
