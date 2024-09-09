package main

import (
	"main/initD"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Title      string
	Lastname   string
	Firstname  string
	Street     string
	Dort       string
	Code       string
	Telephone  string `gorm:"null"`
	Mobil      string
	Email      string `gorm:"unique"`
	Company    string `gorm:"null"`
	Department string `gorm:"null"`
}

func init() {
	initD.LoadEnvVariables()
	initD.ConnectDB()
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
