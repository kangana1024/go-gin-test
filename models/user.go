package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
type Users struct {
	gorm.Model
	Name     string
	Username string
	Password string
}

type CasbinRole struct {
	gorm.Model
	Ptype string
	V0    string
	V1    string
	V2    string
}
