package models

import (
	"final-project-fga/helper"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	ID          uint   `gorm:"primary_key;auto_increment;not_null"`
	Username    string `gorm:"not_null;unique" json:"username" form:"username" validate:"required"`
	Email       string `gorm:"not_null;unique" json:"email" form:"email" validate:"required,email"`
	Password    string `gorm:"not_null" json:"password" form:"password" validate:"required,gte=6"`
	Age         int    `gorm:"not_null" json:"age" form:"age" validate:"required,gte=8"`
	Comment     []Comment
	SocialMedia []SocialMedia
	Photo       []Photo
	Created_At  time.Time
	Updated_At  time.Time
}

func (u *User) BeforeCreate(g *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helper.HashPass(u.Password)
	err = nil
	return
}
