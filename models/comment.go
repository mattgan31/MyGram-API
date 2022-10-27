package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	ID         uint `gorm:"primary_key;auto_increment;not_null"`
	User_ID    uint
	Photo_ID   uint
	Message    string `gorm:"not_null" json:"message" validate:"required"`
	Created_At time.Time
	Updated_At time.Time
	User       *User  //`json:",omitempty"`
	Photo      *Photo //`json:",omitempty"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
