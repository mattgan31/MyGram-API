package models

import "time"

type SocialMedia struct {
	ID               uint   `gorm:"primary_key;auto_increment;not_null"`
	Name             string `gorm:"not_null" json:"name" validate:"required"`
	Social_Media_URL string `gorm:"not_null" json:"social_media_url" validate:"required"`
	User_ID          uint
	Created_At       time.Time
	Updated_At       time.Time
	User             *User
}
