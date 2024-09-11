package controllers

import (
	"fmt"
	"log"
	"main/initD"
	"main/models"

	"github.com/gin-gonic/gin"
)

var err error

func Signup(c *gin.Context) {
	// Get the email/pass off req body

	// Hash the password

	// Create the user

	// Respond

}

func CreateUser(c *gin.Context) {
	var user models.User
	if c.ShouldBind(&user) != nil {
		c.JSON(400, gin.H{"error": "can not bind user"})
		return
	}
	result := initD.DB.Create(&user)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	c.JSON(200, user)
	fmt.Println("User Added", user)

}

func ReadUsers(c *gin.Context) {
	users := []*models.User{}
	result := initD.DB.Find(&users)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	c.JSON(200, users)
	// fmt.Println(result.RowsAffected)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	if c.ShouldBind(&user) != nil {
		c.JSON(400, gin.H{"error": "can not bind user"})
		return
	}
	id := c.Param("id")
	result := initD.DB.Model(&user).Where("ID = ?", id).Updates(user)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {
	var user models.User

	id := c.Param("id")
	result := initD.DB.Delete(&user, id)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	c.JSON(200, nil)
}

func ReadUser(c *gin.Context) {
	user := models.User{}

	id := c.Param("id")
	result := initD.DB.First(&user, id)

	if result.Error != nil {
		log.Fatal(result.Error)
	}
	c.JSON(200, user)
}
