package service

import (
	"context"
	"github.com/cahyacaa/test-julo/internal/app/domain"
)

type Wallet interface {
	InitWallet(ctx context.Context, customerID string) (token string, err error)
	EnableWallet(ctx context.Context, customerID string) (wallet domain.WalletData, err error)
	CheckBalance(ctx context.Context, customerID string) (wallet domain.WalletData, err error)
}
