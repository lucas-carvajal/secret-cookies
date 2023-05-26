package main

import (
	"github.com/gin-gonic/gin"
	"secret-cookies/api"
)

func main() {
	router := gin.Default()

	apiGroup := router.Group("/api")

	apiGroup.POST("/signup", api.Signup)

	router.Run(":8080")
}
