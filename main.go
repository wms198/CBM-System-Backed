package main

import (
	"main/initD"

	"github.com/gin-gonic/gin"
)

func init() {
	initD.LoadEnvVariables()
	initD.ConnectDB()
	initD.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
