package main

import (
	"main/initD"

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
}

func main() {

}
