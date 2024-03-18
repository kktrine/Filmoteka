package model

import (
	"github.com/go-playground/validator"
)

type User struct {
	UserID   int    `gorm:"primaryKey:autoIncrement:true" json:"id"`
	Name     string `gorm:"unique" validate:"required,min=4,max=20" json:"user_name"`
	Password string `validate:"required,min=5,max=20" json:"-"`
	Role     string `json:"role"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) Validate() error {
	validate := validator.New()
	if err := validate.Struct(*u); err != nil {
		return err
	}
	return nil
}
