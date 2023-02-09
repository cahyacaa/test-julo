package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/bsm/redislock"

	"github.com/cahyacaa/test-julo/internal/app/helpers"

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

	err = w.RedisService.SetNX(ctx, helpers.GenerateKey("auth", walletToken), customerID)
	if err != nil {
		return "", err
	}

	if err := w.initWallet(ctx, customerID); err != nil {
		return "", err
	}

	return walletToken, nil
}

func (w *Wallet) EnableWallet(ctx context.Context, customerID string) (wallet domain.WalletData, err error) {
	err = w.RedisService.Get(ctx, helpers.GenerateKey(customerID, domain.Wallet), &wallet)

	if err != nil {
		return
	}

	if !wallet.IsDisabled {
		return wallet, errors.New("Already enabled")
	}

	wallet.IsDisabled = false
	wallet.EnabledAt = time.Now().UTC()

	err = w.RedisService.Set(ctx, helpers.GenerateKey(customerID, domain.Wallet), &wallet)

	return
}

func (w *Wallet) CheckBalance(ctx context.Context, customerID string) (wallet domain.WalletData, err error) {
	err = w.RedisService.Get(ctx, helpers.GenerateKey(customerID, domain.Wallet), &wallet)
	return
}

func (w *Wallet) Deposits(ctx context.Context, customerID, referenceID string, amount float64) (deposits domain.DepositsData, err error) {
	var wallet domain.WalletData
	var depositsData domain.DepositsData
	var lock *redislock.Lock

	// Don't forget to defer Release.
	defer func() {
		errLock := lock.Release(ctx)
		if errLock != nil {
			err = errLock
			return
		}
	}()

	lock, err = w.RedisService.RedisLock.Obtain(ctx, helpers.GenerateKey(referenceID), 100*time.Millisecond, nil)
	if err != nil {
		return
	}

	err = w.RedisService.Get(ctx, helpers.GenerateKey(customerID, domain.Wallet), &wallet)
	if err != nil {
		return
	}

	if wallet.CustomerID == "" {
		return deposits, errors.New("wallet data not found")
	}

	err = w.RedisService.Hget(ctx, helpers.GenerateKey(customerID, domain.Transaction), referenceID, &depositsData)
	if err == nil && depositsData.ID != "" {
		return deposits, errors.New("reference id cannot be same")
	}

	depositsData = domain.DepositsData{
		ID:          customerID,
		DepositedBy: customerID,
		Amount:      amount,
		Type:        domain.Deposits,
		DepositedAt: time.Now().UTC(),
		ReferenceID: referenceID,
	}

	wallet.Balance += amount

	err = w.RedisService.Set(ctx, helpers.GenerateKey(customerID, domain.Wallet), &wallet)
	if err != nil {
		return
	}

	err = w.RedisService.Hset(ctx, helpers.GenerateKey(customerID, domain.Transaction), referenceID, &depositsData)
	if err != nil {
		return
	}

	deposits = domain.DepositsData{
		ID:          wallet.CustomerID,
		DepositedBy: wallet.CustomerID,
		Amount:      amount,
		Type:        domain.Deposits,
		Status:      domain.Success,
		DepositedAt: time.Now(),
		ReferenceID: referenceID,
	}
	return
}

func (w *Wallet) Withdrawals(ctx context.Context, customerID, referenceID string, amount float64) (withdrawals domain.WithdrawalsData, err error) {
	var wallet domain.WalletData
	var depositsData domain.WithdrawalsData
	var lock *redislock.Lock

	// Don't forget to defer Release.
	defer func() {
		errLock := lock.Release(ctx)
		if errLock != nil {
			err = errLock
			return
		}
	}()

	lock, err = w.RedisService.RedisLock.Obtain(ctx, helpers.GenerateKey(referenceID), 100*time.Millisecond, nil)
	if err != nil {
		return
	}

	err = w.RedisService.Get(ctx, helpers.GenerateKey(customerID, domain.Wallet), &wallet)
	if err != nil {
		return
	}

	if wallet.CustomerID == "" {
		return withdrawals, errors.New("withdrawals data not found")
	}

	err = w.RedisService.Hget(ctx, helpers.GenerateKey(customerID, domain.Transaction), referenceID, &depositsData)
	if err == nil && depositsData.ID != "" {
		return withdrawals, errors.New("reference id cannot be same")
	}

	depositsData = domain.WithdrawalsData{
		ID:          customerID,
		WithdrawnBy: customerID,
		Amount:      amount,
		Type:        domain.Withdrawal,
		WithdrawnAt: time.Now().UTC(),
		ReferenceID: referenceID,
	}

	wallet.Balance -= amount

	if wallet.Balance < 0 {
		err = errors.New("insufficient balance")
		return
	}
	err = w.RedisService.Set(ctx, helpers.GenerateKey(customerID, domain.Wallet), &wallet)
	if err != nil {
		return
	}

	err = w.RedisService.Hset(ctx, helpers.GenerateKey(customerID, domain.Transaction), referenceID, &depositsData)
	if err != nil {
		return
	}

	withdrawals = domain.WithdrawalsData{
		ID:          wallet.CustomerID,
		WithdrawnBy: wallet.CustomerID,
		Amount:      amount,
		Type:        domain.Withdrawal,
		Status:      domain.Success,
		WithdrawnAt: time.Now(),
		ReferenceID: referenceID,
	}
	return
}

func (w *Wallet) initWallet(ctx context.Context, customerID string) (err error) {
	err = w.RedisService.SetNX(ctx, helpers.GenerateKey(customerID, domain.Wallet), domain.WalletData{
		CustomerID: customerID,
		Balance:    0,
		IsDisabled: false,
		EnabledAt:  time.Now().UTC(),
	})

	return
}
