package models

import (
	"gorm.io/gorm"
)

type User struct {
	Name     string `gorm:"column:name;size:255;not null;"`
	Email    string `gorm:"column:email;unique;not null;index"`
	Password string `gorm:"column:password;not null"`
	gorm.Model
}
