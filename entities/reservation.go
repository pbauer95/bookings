package entities

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	User      User
	UserID    uint
	Room      Room
	RoomID    uint
	Phone     sql.NullString
	StartDate time.Time
	EndDate   time.Time
}
