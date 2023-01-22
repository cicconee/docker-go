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

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		log.Println("no port provided... defaulting to 5000")
		httpPort = "5000"
	}

	r.Run("0.0.0.0:" + httpPort)
}