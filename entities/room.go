package entities

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	RoomName string
}
