package middleware

import (
	"context"
	"github.com/cahyacaa/test-julo/internal/app/domain"
	"github.com/cahyacaa/test-julo/internal/app/pkg/error"
	"github.com/cahyacaa/test-julo/internal/app/pkg/redis"
	responseFormat "github.com/cahyacaa/test-julo/internal/app/pkg/response_format"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckWalletStatusHandler(ctx context.Context, redisService redis.RedisService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var customerID string
		var wallet domain.WalletData

		if value, ok := c.Get("customer_id"); !ok {
			response := responseFormat.HandleError[error.Format]("token is empty", http.StatusBadRequest)
			c.AbortWithStatusJSON(response.StatusCode, response)
			return
		} else {
			customerID = value.(string)
		}

		if err := redisService.Get(ctx, customerID, &wallet); err != nil {
			response := responseFormat.HandleError("customer not found", http.StatusUnauthorized)
			c.AbortWithStatusJSON(response.StatusCode, response)
			return
		}

		if wallet.IsDisabled {
			response := responseFormat.HandleError("wallet is disabled", http.StatusUnauthorized)
			c.AbortWithStatusJSON(response.StatusCode, response)
			return
		}

		c.Set("wallet_data", wallet)
		c.Next()
	}
}
