package controllers

import (
	"fmt"
	"log"
	"main/initD"
	"main/models"

	"github.com/gin-gonic/gin"
)

func CreateAppointment(c *gin.Context) {
	var appiontment models.Appointment
	if c.ShouldBind(&appiontment) != nil {
		c.JSON(400, gin.H{"error": "can not bind appointmenr"})
		return
	}
	result := initD.DB.Create(&appiontment)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}
	c.JSON(200, appiontment)
	fmt.Println("Appointment added", appiontment)
}

func ReadAppointments(c *gin.Context) {
	appointements := []*models.Appointment{}
	result := initD.DB.Find(&appointements)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	c.JSON(200, appointements)
}

func UpdateAppointment(c *gin.Context) {
	var appointment models.Appointment
	if c.ShouldBind(&appointment) != nil {
		c.JSON(400, gin.H{"error": "can not bind user"})
		return
	}
	id := c.Param("id")
	result := initD.DB.Model(&appointment).Where("ID = ?", id).Updates(appointment)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	c.JSON(200, appointment)
}

func DeleteAppointment(c *gin.Context) {
	var appiontment models.Appointment

	id := c.Param("id")
	result := initD.DB.Delete(&appiontment, id)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	c.JSON(200, nil)
}

func Readappointment(c *gin.Context) {
	var appointement models.Appointment
	id := c.Param("id")
	result := initD.DB.First(&appointement, id)

	if result.Error != nil {
		log.Fatal(result.Error)
	}
	c.JSON(200, appointement)

}
