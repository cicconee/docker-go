package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world!")

	r := gin.New()

	r.GET("/", func(ctx *gin.Context) {
		log.Printf("request from %s\n", ctx.RemoteIP())
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
			"your_ip": ctx.RemoteIP(),
		})
	})

	r.GET("/sir", func(ctx *gin.Context) {
		requestIP := ctx.Request.Header["X-Forwarded-For"]
		log.Printf("/sir: request from %s\n", requestIP)

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello, SIR!",
		})
	})

	r.GET("/headers", func(ctx *gin.Context) {
		headers := ctx.Request.Header

		ctx.JSON(http.StatusOK, headers)
	})

	r.GET("/a", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "okay at /a")
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		log.Println("no port provided... defaulting to 5000")
		httpPort = "5000"
	}

	r.Run(":" + httpPort)
}
