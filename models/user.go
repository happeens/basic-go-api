package models

import (
	_ "github.com/jinzhu/gorm"
	"time"
)

type User struct {
	ID        uint   `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name      string `json:"description" gorm:"not null;unique"`
	Email     string `json:"email" gorm:"not null;unique"`
	Password  string `json:"-" gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `json:"-"`
}
