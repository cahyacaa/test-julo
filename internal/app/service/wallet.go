package service

import (
	"context"

	"github.com/cahyacaa/test-julo/internal/app/domain"
)

type Wallet interface {
	InitWallet(ctx context.Context, customerID string) (token string, err error)
	EnableWallet(ctx context.Context, customerID string) (wallet domain.WalletData, err error)
	CheckBalance(ctx context.Context, customerID string) (wallet domain.WalletData, err error)
	Deposits(ctx context.Context, customerID, referenceID string, amount float64) (wallet domain.DepositsData, err error)
	Withdrawals(ctx context.Context, customerID, referenceID string, amount float64) (wallet domain.WithdrawalsData, err error)
	ViewTransactions(ctx context.Context, customerID string) (wallet []domain.TransactionData, err error)
}
