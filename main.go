package main

import (
	"main/controllers"
	"main/initD"

	"github.com/gin-gonic/gin"
)

func init() {
	initD.LoadEnvVariables()
	initD.ConnectDB()
	initD.SyncDatabase()
}

// GET retrieves data.
// POST creates data.
// PUT updates data entirely.
// PATCH allows partially updating data.
// DELETE removes data.
func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/createUser", controllers.CreateUser)

	r.GET("/users", controllers.ReadUsers)
	r.GET("/user/:id", controllers.ReadUser)
	r.PATCH("/user/:id", controllers.UpdateUser)
	r.DELETE("/user/:id", controllers.DeleteUser)

	r.Run()
}
