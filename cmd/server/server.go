package server

import (
	"fmt"
	"log"
	"os"
	"time"

	"dockerize/cmd/middlewares"
	"github.com/gin-gonic/gin"
)

// Start create routes, set port and run app.
func Start() {
	router := createRouter()
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	if err := router.Run(":" + port); err != nil {
		fmt.Printf("%s", err)
	}
}

func createRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] %s %s %d \"%s\" %s",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC822),
			param.Method,
			param.Path,
			param.StatusCode,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())
	router.Use(middlewares.ErrorHandler)
	setRoutes(router)

	return router
}
