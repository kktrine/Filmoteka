package model

import (
	"github.com/go-playground/validator"
)

type Film struct {
	Id           int     `gorm:"primaryKey:autoIncrement:true" json:"-"`
	Title        string  `gorm:"unique" json:"title" validate:"required,min=1,max=150"`
	Description  string  `gorm:"description" json:"description" validate:"omitempty,max=10000"`
	DatePremiere int     `json:"date_of_premiere" validate:"omitempty"`
	Rating       float64 `json:"rating" validate:"required,min=0,max=10"`
	Actors       []Actor `gorm:"many2many:actors_films;primaryKey:false;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func (Film) TableName() string {
	return "films"
}

func (f *Film) Validate() error {
	validate := validator.New()
	if err := validate.Struct(*f); err != nil {
		return err
	}
	return nil
}

func (f *Film) ValidateTitle() bool {
	validate := validator.New()
	if err := validate.StructPartial(f, "Title"); err != nil {
		return false
	}
	return true
}

func (f *Film) ValidateDescription() bool {
	validate := validator.New()
	if err := validate.StructPartial(f, "Description"); err != nil {
		return false
	}
	return true
}

func (f *Film) ValidateDatePremiere() bool {
	validate := validator.New()
	if err := validate.StructPartial(f, "DatePremiere"); err != nil {
		return false
	}
	return true
}

func (f *Film) ValidateRating() bool {
	validate := validator.New()
	if err := validate.StructPartial(f, "Rating"); err != nil {
		return false
	}
	return true
}
