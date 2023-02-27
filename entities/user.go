package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName   string
	LastName    string
	Email       string `gorm:"unique"`
	Password    string
	AccessLevel int
}
