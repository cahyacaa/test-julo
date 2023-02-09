package controller

import (
	"context"

	"github.com/cahyacaa/test-julo/internal/app/handler"
	"github.com/cahyacaa/test-julo/internal/app/pkg/middleware"
	"github.com/cahyacaa/test-julo/internal/app/pkg/redis"
	"github.com/cahyacaa/test-julo/internal/app/usecase"
	"github.com/gin-gonic/gin"
)

type Dependency struct {
	RedisService redis.RedisService
}

func Router(ctx context.Context, r *gin.Engine, dep Dependency) *gin.Engine {

	//init dependency

	walletUcase := usecase.NewWalletUcase(dep.RedisService)
	walletController := handler.NewWallerHandler(walletUcase)

	//router non middleware
	r.POST("/api/v1/init", walletController.InitWalletAccount)

	//init wallet router
	walletRouter := r.Group("/api/v1")

	// middleware for wallet router group
	walletRouter.Use(middleware.Authorization(ctx, dep.RedisService))
	walletRouter.POST("/wallet", walletController.EnableWallet)
	walletRouter.PATCH("/wallet")

	walletFeatureRouter := walletRouter.Group("")
	walletFeatureRouter.Use(middleware.CheckWalletStatusHandler(ctx, dep.RedisService))
	walletFeatureRouter.GET("/wallet", walletController.CheckBalance)
	walletFeatureRouter.GET("/wallet/transactions", walletController.ViewTransactions)
	walletFeatureRouter.POST("/wallet/deposits", walletController.Deposits)
	walletFeatureRouter.POST("/wallet/withdrawals", walletController.Withdrawals)

	return r
}
