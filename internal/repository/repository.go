package Repo

import (
	"errors"

	"github.com/pbauer95/bookings/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Repo struct {
	Connection *gorm.DB
}

func InitializeDb() (*Repo, error) {
	conn, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return &Repo{}, err
	}

	db := Repo{
		Connection: conn,
	}

	// Migrate the schema
	db.Connection.AutoMigrate(
		&entities.User{},
		&entities.Room{},
		&entities.Reservation{},
		&entities.Restriction{},
		&entities.RoomRestriction{},
	)
	return &db, nil
}

func (db *Repo) SeedRepo() error {
	if db.Connection == nil {
		return errors.New("Repo has not been initialized")
	}

	db.Connection.Create(&entities.User{
		FirstName:   "Peter",
		LastName:    "Bauer",
		Email:       "peter.bauer@capgemini.com",
		AccessLevel: 0,
		Password:    "",
	})

	return nil
}
