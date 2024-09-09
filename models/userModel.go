package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Appointments []Appointment
	Title        string
	Lastname     string
	Firstname    string
	Street       string
	Dort         string
	Code         string
	Telephone    string `gorm:"null"`
	Mobil        string
	Email        string `gorm:"unique"`
	Password     string
	Company      string `gorm:"null"`
	Department   string `gorm:"null"`
}
