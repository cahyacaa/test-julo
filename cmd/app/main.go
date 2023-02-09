package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/cahyacaa/test-julo/cmd/config"
	"github.com/cahyacaa/test-julo/internal/app/controller"
	"github.com/cahyacaa/test-julo/internal/app/pkg/middleware"
	"github.com/cahyacaa/test-julo/internal/app/pkg/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	//init context
	ctx := context.Background()

	//init global config
	config.InitializeConfig()

	// init router
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(middleware.CORSMiddleware())

	//init redis
	redisService := redis.NewRedisService()

	err := redisService.InitRedis(config.GlobalConfig.Cache)
	if err != nil {
		log.Fatal(err)
	}

	//health check routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	//wallet router
	r = controller.Router(ctx, r, controller.Dependency{
		RedisService: redisService,
	})

	err = r.Run(fmt.Sprintf(":%s", config.GlobalConfig.App.Port))
	if err != nil {
		log.Fatal(err)
	}
}
