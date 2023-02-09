package usecase

import (
	"context"
	"errors"
	"github.com/cahyacaa/test-julo/internal/app/domain"
	"github.com/cahyacaa/test-julo/internal/app/pkg/redis"
	"time"
)

const walletToken = "cb04f9f26632ad602f14acef21c58f58f6fe5fb55a"

type Wallet struct {
	RedisService redis.RedisService
}

func NewWalletUcase(redisService redis.RedisService) Wallet {
	return Wallet{
		RedisService: redisService,
	}
}

func (w *Wallet) InitWallet(ctx context.Context, customerID string) (token string, err error) {
	if customerID == "" {
		err = errors.New("missing data for required field.")
		return
	}

	err = w.RedisService.SetNX(ctx, walletToken, customerID)
	if err != nil {
		return "", err
	}

	if err := w.initWallet(ctx, customerID); err != nil {
		return "", err
	}

	return walletToken, nil
}

func (w *Wallet) EnableWallet(ctx context.Context, customerID string) (walletData domain.WalletData, err error) {
	err = w.RedisService.Get(ctx, customerID, &walletData)

	if err != nil {
		return
	}

	if !walletData.IsDisabled {
		return walletData, errors.New("Already enabled")
	}

	err = w.RedisService.Set(ctx, customerID, &domain.WalletData{
		IsDisabled: false,
	})

	return
}

func (w *Wallet) CheckBalance(ctx context.Context, customerID string) (wallet domain.WalletData, err error) {
	err = w.RedisService.Get(ctx, customerID, &wallet)
	return
}

func (w *Wallet) initWallet(ctx context.Context, customerID string) (err error) {

	err = w.RedisService.SetNX(ctx, customerID, domain.WalletData{
		CustomerID: customerID,
		Balance:    0,
		IsDisabled: false,
		EnabledAt:  time.Now().UTC(),
	})

	return
}
