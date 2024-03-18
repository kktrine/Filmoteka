package repository

import (
	// "filmoteka_server/api/models"
	"filmoteka_server/internal/model"
	"filmoteka_server/models"
)

type Repository interface {
	UserOperations
	FilmOperations
	ActorOperations
	Stop() error
}

type UserOperations interface {
	Login(userName, password string) (string, error)
}

type FilmOperations interface {
	InsertFilm(model.Film, []model.Actor) error
	UpdateFilm(model.Film, model.Film) error
	DeleteFilm(string, int) error
	SelectFilmsPattern(string, string) ([]model.Film, error)
	SelectSortedFilms(string) ([]model.Film, error)
}

type ActorOperations interface {
	InsertActor(model.Actor) error
	UpdateActor(model.Actor, model.Actor) error
	DeleteActor(model.Actor) error
	SelectAllActors() ([]models.ActorFilm, error)
}
