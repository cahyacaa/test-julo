package usecase

import (
	"context"
	"errors"
	"github.com/cahyacaa/test-julo/internal/app/domain"
	"github.com/cahyacaa/test-julo/internal/app/pkg/redis"
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

	err = w.RedisService.Set(ctx, customerID, &domain.WalletAuth{
		Token:      walletToken,
		IsDisabled: false,
	})

	if err != nil {
		return "", err
	}

	return walletToken, nil
}

func (w *Wallet) EnableWallet(ctx context.Context, customerID string) (err error) {
	return
}
