package models

import "time"

type Photo struct {
	ID         uint   `gorm:"primary_key;auto_increment;not_null"`
	Title      string `gorm:"not_null" json:"title" validate:"required"`
	Caption    string `json:"caption"`
	Photo_URL  string `gorm:"not_null" json:"photo_url" validate:"required"`
	User_ID    uint
	User       *User
	Comment    *Comment
	Created_At time.Time
	Updated_At time.Time
}
