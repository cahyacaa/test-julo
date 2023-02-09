package middleware

import (
	"context"
	"errors"
	"github.com/cahyacaa/test-julo/internal/app/domain"
	"github.com/cahyacaa/test-julo/internal/app/pkg/redis"
	responseFormat "github.com/cahyacaa/test-julo/internal/app/pkg/response_format"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckWalletStatusHandler(ctx context.Context, redisService redis.RedisService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var walletAuth domain.WalletAuth
		var customerID string

		if value, ok := c.Get("customerID"); ok {
			customerID = value.(string)
		}

		if err := redisService.Get(ctx, customerID, &walletAuth); err != nil {
			response := responseFormat.HandleError(err, http.StatusBadRequest)
			c.AbortWithStatusJSON(response.StatusCode, response)
			return
		}

		if walletAuth.IsDisabled {
			response := responseFormat.HandleError(errors.New("wallet is disabled"), http.StatusUnauthorized)
			c.AbortWithStatusJSON(response.StatusCode, response)
			return
		}
		c.Next()
	}
}
