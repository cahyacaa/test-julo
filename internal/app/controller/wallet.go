package controller

import (
	"context"
	"github.com/cahyacaa/test-julo/internal/app/handler"
	"github.com/cahyacaa/test-julo/internal/app/pkg/middleware"
	"github.com/cahyacaa/test-julo/internal/app/pkg/redis"
	"github.com/cahyacaa/test-julo/internal/app/usecase"
	"github.com/gin-gonic/gin"
)

func Router(ctx context.Context, r *gin.Engine) *gin.Engine {

	//init dependency
	redisService := redis.NewRedisService()
	walletUcase := usecase.NewWalletUcase(redisService)
	walletController := handler.NewWallerHandler(walletUcase)

	//router non middleware
	r.POST("/api/v1/init", walletController.InitWalletAccount)

	//init wallet router
	walletRouter := r.Group("/api/v1")

	// middleware for wallet router group
	walletRouter.Use(middleware.Authorization(ctx, redisService))
	walletRouter.POST("/wallet", walletController.EnableWallet)
	walletRouter.PATCH("/wallet")

	walletFeatureRouter := walletRouter.Group("")
	walletFeatureRouter.Use(middleware.CheckWalletStatusHandler(ctx, redisService))
	walletFeatureRouter.GET("/wallet", walletController.CheckBalance)

	return r
}
