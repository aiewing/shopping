package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"shopping/api"
	"shopping/utils/graceful"
	"time"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	engine := gin.Default()
	registerMiddlewares(engine)

	api.RegisterHandlers(engine)
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	srv := &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()
	graceful.ShutdownGin(srv, time.Second*5)
}

// 注册中间件
func registerMiddlewares(engine *gin.Engine) {
	engine.Use(
		gin.LoggerWithFormatter(
			func(params gin.LogFormatterParams) string {
				return fmt.Sprintf(
					"%s - [%s] \"%s %s %s %d %s %s\"\n",
					params.ClientIP,
					params.TimeStamp.Format(time.RFC3339),
					params.Method,
					params.Path,
					params.Request.Proto,
					params.StatusCode,
					params.Latency,
					params.ErrorMessage,
				)
			}))
	engine.Use(gin.Recovery())
}
