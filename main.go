package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thuongthanhto/go-shop-backend/models"
)

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "Hello world!"})
	})

	models.ConnectDatabase()

	r.Run()
}
