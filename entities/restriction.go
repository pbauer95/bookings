package entities

import "gorm.io/gorm"

type Restriction struct {
	gorm.Model
	Name string
}
