package models

import (
	"time"

	"gorm.io/gorm"
)

// User has many appointments, UserId is the foreign key
type Appointment struct {
	gorm.Model
	UserId      uint
	Date        time.Time
	Explanation string
}
