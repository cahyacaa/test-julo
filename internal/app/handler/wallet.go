package handler

import (
	"net/http"
	"time"

	"github.com/cahyacaa/test-julo/internal/app/domain"
	responseFormat "github.com/cahyacaa/test-julo/internal/app/pkg/response_format"
	"github.com/cahyacaa/test-julo/internal/app/usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Wallet struct {
	WalletUsecase usecase.Wallet
}

func NewWallerHandler(wUcase usecase.Wallet) Wallet {
	return Wallet{
		WalletUsecase: wUcase,
	}
}

func (w *Wallet) InitWalletAccount(c *gin.Context) {
	var req domain.InitWalletAccountRequest
	err := c.Bind(&req)
	if err != nil {
		l := err.(validator.ValidationErrors)[0].Field()
		response := responseFormat.HandleError(l, http.StatusBadRequest)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	token, err := w.WalletUsecase.InitWallet(c, req.CustomerXID)

	if err != nil {
		response := responseFormat.HandleError(err.Error(), http.StatusInternalServerError)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	//convert to response field
	data := domain.InitWalletAccountResponse{
		Token: token,
	}

	response := responseFormat.HandleSuccess[domain.InitWalletAccountResponse](data)
	c.JSON(response.StatusCode, response)
}

func (w *Wallet) CheckBalance(c *gin.Context) {

	var walletData domain.WalletData

	if val, ok := c.Get("wallet_data"); ok {
		walletData = val.(domain.WalletData)
	}

	wallet, err := w.WalletUsecase.CheckBalance(c, walletData.CustomerID)

	if err != nil {
		response := responseFormat.HandleError(err.Error(), http.StatusInternalServerError)
		c.JSON(response.StatusCode, response)
		return
	}

	//convert to response field
	data := domain.WalletResponse{
		Wallet: struct {
			ID        string              `json:"id"`
			OwnedBy   string              `json:"owned_by"`
			Balance   float64             `json:"balance"`
			Status    domain.WalletStatus `json:"status"`
			EnabledAt time.Time           `json:"enabled_at"`
		}{
			ID:        wallet.CustomerID,
			OwnedBy:   wallet.CustomerID,
			Balance:   wallet.Balance,
			EnabledAt: wallet.EnabledAt,
		},
	}

	if wallet.IsDisabled {
		data.Wallet.Status = domain.Disabled
	} else {
		data.Wallet.Status = domain.Enabled
	}
	response := responseFormat.HandleSuccess[domain.WalletResponse](data)
	c.JSON(response.StatusCode, response)
}

func (w *Wallet) EnableWallet(c *gin.Context) {
	var customerID string

	if val, ok := c.Get("customer_id"); ok {
		customerID = val.(string)
	}

	wallet, err := w.WalletUsecase.EnableWallet(c, customerID)

	if err != nil {
		response := responseFormat.HandleError(err.Error(), http.StatusBadRequest)
		c.JSON(response.StatusCode, response)
		return
	}

	//convert to response field
	data := domain.WalletResponse{
		Wallet: struct {
			ID        string              `json:"id"`
			OwnedBy   string              `json:"owned_by"`
			Balance   float64             `json:"balance"`
			Status    domain.WalletStatus `json:"status"`
			EnabledAt time.Time           `json:"enabled_at"`
		}{
			ID:        wallet.CustomerID,
			OwnedBy:   wallet.CustomerID,
			Balance:   wallet.Balance,
			EnabledAt: wallet.EnabledAt,
		},
	}
	if wallet.IsDisabled {
		data.Wallet.Status = domain.Disabled
	} else {
		data.Wallet.Status = domain.Enabled
	}
	response := responseFormat.HandleSuccess[domain.WalletResponse](data)
	c.JSON(response.StatusCode, response)
}

func (w *Wallet) Deposits(c *gin.Context) {
	var req domain.DepositsRequest
	var customerID string

	err := c.Bind(&req)
	if err != nil {
		l := err.(validator.ValidationErrors)[0].Field()
		response := responseFormat.HandleError(l, http.StatusBadRequest)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if val, ok := c.Get("customer_id"); ok {
		customerID = val.(string)
	}

	deposits, err := w.WalletUsecase.Deposits(c, customerID, req.ReferenceID, req.Amount)
	if err != nil {
		response := responseFormat.HandleError(err.Error(), http.StatusBadRequest)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	data := domain.DepositsResponse{Deposits: deposits}
	response := responseFormat.HandleSuccess[domain.DepositsResponse](data)
	c.JSON(response.StatusCode, response)
}

func (w *Wallet) Withdrawals(c *gin.Context) {
	var req domain.WithdrawalsRequest
	var customerID string

	err := c.Bind(&req)
	if err != nil {
		l := err.(validator.ValidationErrors)[0].Field()
		response := responseFormat.HandleError(l, http.StatusBadRequest)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if val, ok := c.Get("customer_id"); ok {
		customerID = val.(string)
	}

	withdrawal, err := w.WalletUsecase.Withdrawals(c, customerID, req.ReferenceID, req.Amount)
	if err != nil {
		response := responseFormat.HandleError(err.Error(), http.StatusBadRequest)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	data := domain.WithdrawalsResponse{Withdrawal: withdrawal}
	response := responseFormat.HandleSuccess[domain.WithdrawalsResponse](data)
	c.JSON(response.StatusCode, response)

}
