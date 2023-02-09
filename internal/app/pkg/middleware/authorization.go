package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/cahyacaa/test-julo/internal/app/helpers"
	"github.com/cahyacaa/test-julo/internal/app/pkg/redis"
	responseFormat "github.com/cahyacaa/test-julo/internal/app/pkg/response_format"
	"github.com/gin-gonic/gin"
)

func Authorization(ctx context.Context, redisService redis.RedisService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var customerID string

		authHeader := c.Request.Header.Get("Authorization")
		if len(authHeader) == 0 {
			response := responseFormat.HandleError("authorization is empty", http.StatusUnauthorized)
			c.AbortWithStatusJSON(response.StatusCode, response)
			return
		}

		tokenInfo := strings.Fields(authHeader)
		if len(tokenInfo) < 2 {
			response := responseFormat.HandleError("token is invalid", http.StatusUnauthorized)
			c.AbortWithStatusJSON(response.StatusCode, response)
			return
		}

		if strings.ToLower(tokenInfo[0]) != "token" {
			response := responseFormat.HandleError("not using Token as a bearer", http.StatusUnauthorized)
			c.AbortWithStatusJSON(response.StatusCode, response)
			return
		}

		if len(tokenInfo) == 0 {
			response := responseFormat.HandleError("token is invalid", http.StatusUnauthorized)
			c.AbortWithStatusJSON(response.StatusCode, response)
			return
		}

		if err := redisService.Get(ctx, helpers.GenerateKey("auth", tokenInfo[1]), &customerID); err != nil {
			response := responseFormat.HandleError("token not found", http.StatusUnauthorized)
			c.AbortWithStatusJSON(response.StatusCode, response)
			return
		}

		c.Set("customer_id", customerID)
		c.Next()
	}
}
