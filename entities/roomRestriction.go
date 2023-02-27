package entities

import (
	"time"

	"gorm.io/gorm"
)

type RoomRestriction struct {
	gorm.Model
	Reservation   Reservation
	ReservationID uint `gorm:"index"`
	Room          Room
	RoomID        uint `gorm:"index"`
	Restriction   Restriction
	RestrictionID uint
	StartDate     time.Time `gorm:"index:start_end"`
	EndDate       time.Time `gorm:"index:start_end"`
}
