package models

import (
	"gorm.io/gorm"
)

// type Contact struct {
// 	gorm.Model
// 	ContactID   int    `gorm:"primary_key" json:id`
// 	Name        string `json:"name"`
// 	PhoneNumber string `json:"phone_number"`
// 	userID      int    `json:"userId"`
// 	User        User   `json:"-"`
// }

type Contact struct {
	gorm.Model
	Name        string
	PhoneNumber string `gorm:"unique_index`
	Email       string `gorm:"typevarchar(100);unique_index`
}
