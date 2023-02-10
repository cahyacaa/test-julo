package handler

import (
	"github.com/cahyacaa/test-julo/internal/app/domain"
	responseFormat "github.com/cahyacaa/test-julo/internal/app/pkg/response_format"
	"github.com/cahyacaa/test-julo/internal/app/usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
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
		Wallet: domain.WalletResponseData{
			ID:      wallet.CustomerID,
			OwnedBy: wallet.CustomerID,
			Balance: wallet.Balance,
		},
	}

	if !wallet.IsDisabled {
		data.Wallet.Status = domain.Enabled
		data.Wallet.EnabledAt = &wallet.EnabledAt
		data.Wallet.DisabledAt = nil
	} else {
		data.Wallet.Status = domain.Disabled
		data.Wallet.DisabledAt = &wallet.DisabledAt
		data.Wallet.EnabledAt = nil
	}
	response := responseFormat.HandleSuccess[domain.WalletResponse](data)
	c.JSON(response.StatusCode, response)
}

func (w *Wallet) ViewTransactions(c *gin.Context) {
	var customerID string

	if val, ok := c.Get("customer_id"); ok {
		customerID = val.(string)
	}

	transactions, err := w.WalletUsecase.ViewTransactions(c, customerID)

	if err != nil {
		response := responseFormat.HandleError(err.Error(), http.StatusInternalServerError)
		c.JSON(response.StatusCode, response)
		return
	}

	//convert to response field
	data := domain.TransactionsResponse{
		Transactions: transactions,
	}

	response := responseFormat.HandleSuccess[domain.TransactionsResponse](data)
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
		Wallet: domain.WalletResponseData{
			ID:      wallet.ID,
			OwnedBy: wallet.OwnedBy,
			Balance: wallet.Balance,
			Status:  wallet.Status,
		},
	}

	if wallet.Status == domain.Enabled {
		data.Wallet.EnabledAt = &wallet.EnabledAt
		data.Wallet.DisabledAt = nil
	}

	response := responseFormat.HandleSuccess[domain.WalletResponse](data)
	c.JSON(response.StatusCode, response)
}

func (w *Wallet) DisableWallet(c *gin.Context) {
	var customerID string

	if val, ok := c.Get("customer_id"); ok {
		customerID = val.(string)
	}

	wallet, err := w.WalletUsecase.DisableWallet(c, customerID)

	if err != nil {
		response := responseFormat.HandleError(err.Error(), http.StatusBadRequest)
		c.JSON(response.StatusCode, response)
		return
	}

	//convert to response field
	data := domain.WalletResponse{
		Wallet: domain.WalletResponseData{
			ID:      wallet.ID,
			OwnedBy: wallet.OwnedBy,
			Balance: wallet.Balance,
			Status:  wallet.Status,
		},
	}

	if wallet.Status == domain.Disabled {
		data.Wallet.DisabledAt = &wallet.DisabledAt
		data.Wallet.EnabledAt = nil
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
