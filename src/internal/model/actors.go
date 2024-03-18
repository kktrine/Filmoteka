package model

import (
	"time"

	"github.com/go-playground/validator"
)

type Actor struct {
	Id           int       `gorm:"primaryKey:autoIncrement:true" json:"-"`
	Name         string    `json:"name" validate:"required,min=1"`
	Gender       string    `json:"gender" validate:"min=1"`
	DateBirthday time.Time `gorm:"type:date" json:"date_birthday" validate:"required,min=1"`
	Films        []Film    `gorm:"many2many:actors_films;primaryKey:false;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

type Actors_films struct {
	ActorID int
	FilmID  int
}

func (Actor) TableName() string {
	return "actor"
}

func (a *Actor) Validate() error {
	validate := validator.New()
	if err := validate.Struct(*a); err != nil {
		return err
	}
	return nil
}

func (a *Actor) ValidateName() bool {
	validate := validator.New()
	if err := validate.StructPartial(a, "Name"); err != nil {
		return false
	}
	return true
}

func (a *Actor) ValidateBirthday() bool {
	validate := validator.New()
	if err := validate.StructPartial(a, "date_birthday"); err != nil {
		return false
	}
	return true
}

func (a *Actor) ValidateGender() bool {
	validate := validator.New()
	if err := validate.StructPartial(a, "Gender"); err != nil {
		return false
	}
	return true
}
