package initD

import (
	"main/models"
)

// Pass struct and using AutoMigrate to create tables
func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Appointment{})

}
