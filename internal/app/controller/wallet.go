package controller

import (
	"context"
	"github.com/cahyacaa/test-julo/internal/app/handler"
	"github.com/cahyacaa/test-julo/internal/app/pkg/middleware"
	"github.com/cahyacaa/test-julo/internal/app/pkg/redis"
	"github.com/cahyacaa/test-julo/internal/app/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router(ctx context.Context, r *gin.Engine) *gin.Engine {

	//init dependency
	redisService := redis.NewRedisService()
	walletUcase := usecase.NewWalletUcase(redisService)
	walletController := handler.NewWallerHandler(walletUcase)

	//router non middleware
	r.POST("/init", walletController.InitWalletAccount)

	//init wallet router
	walletRouter := r.Group("/api/v1")

	// middleware for wallet router group
	walletRouter.Use(middleware.CheckWalletStatusHandler(ctx, redisService))
	walletRouter.GET("/wallet", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	return r
}
