package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)


type Product struct {
	GormModel
	Title string `json:"title" form:"title" valid:"required-Title of your product is required"`
	Description string `json:"description" form:"description" valid:"required-Description of your product is required"`
	UserID uint
	User *User
}

func (u *Product) BeforeCreate(tx *gorm.DB) (err error){
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil{
		err = errCreate
		return 
	}

	err = nil
	return
}

func (u *Product) BeforeUpdate(tx *gorm.DB) (err error){
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil{
		err = errCreate
		return 
	}

	err = nil
	return
}