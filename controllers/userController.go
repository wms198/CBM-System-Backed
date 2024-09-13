package controllers

import (
	"fmt"
	"log"
	"main/initD"
	"main/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// Get the email/pass off req body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	// Hash the password
	hash, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	// Create the user
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initD.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}
	// Respond
	c.JSON(http.StatusOK, gin.H{})

}

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	var user models.User
	initD.DB.First(&user, "email=?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invaild email",
		})
		return
	}

	// Compare sent in pass with saved user pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invail  password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}
	// send it back as token
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})

}

func Validate(c *gin.Context) {
	println(c.Request.Header)
	c.JSON(http.StatusOK, gin.H{
		"message": "I'm logged in",
	})
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
