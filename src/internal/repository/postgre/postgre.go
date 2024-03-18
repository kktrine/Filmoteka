package postgre

import (
	"errors"
	// "filmoteka_server/api/models"
	"filmoteka_server/internal/config"
	"filmoteka_server/internal/model"
	"filmoteka_server/models"

	//"filmoteka_server/models"
	"fmt"

	"github.com/go-openapi/strfmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	Db *gorm.DB
}

func NewPostgresRepository(cfg config.DbConfig) *Postgres {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", cfg.Host, cfg.User, cfg.Password, cfg.DbName, cfg.Port)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic("coudn't connect to database")
	}
	if !db.Migrator().HasTable(&model.Actor{}) {
		db.AutoMigrate(&model.User{}, &model.Actor{}, &model.Film{})
	}
	return &Postgres{db}
}

func (p *Postgres) Stop() error {
	val, err := p.Db.DB()
	if err != nil {
		return errors.New("failed to get database error: " + err.Error())
	}
	return val.Close()
}

func (p *Postgres) Login(userName, password string) (string, error) {
	const op = "repository.postgre.Login"
	var user model.User
	result := p.Db.Where("name = ?", userName).Take(&user)
	var err error = nil
	if result.Error != nil {
		err = errors.New("user not found")
	} else {
		if user.Password == password {
			return user.Role, nil
		} else {
			err = errors.New("incorrect pasword or login")
		}
	}
	return "", check(op, err)
}

func (p *Postgres) InsertActor(actor model.Actor) error {
	const op = "repository.postgre.InsertActor"
	var actors model.Actor
	p.Db.Where("name = ? AND date_birthday = ?", actor.Name, actor.DateBirthday).First(&actors)
	if actors.ValidateName() {
		return check(op, errors.New("this actor already exist"))
	}
	result := p.Db.Create(&actor)
	var err error
	if result.Error != nil {
		err = errors.New("faile to create actor")
	}
	return check(op, err)
}

func (p *Postgres) DeleteActor(actor model.Actor) error {
	const op = "repository.postgre.DeleteActor"
	var record model.Actor
	result := p.Db.Model(model.Actor{}).Where("name = ? AND date_birthday = ?", actor.Name, actor.DateBirthday).Delete(record)
	return check(op, result.Error)
}

func (p *Postgres) UpdateActor(actorOld, actor model.Actor) error {
	const op = "repository.postgre.UpdateActor"
	result := p.Db.Model(model.Actor{}).Where("name = ? AND date_birthday = ?", actorOld.Name, actorOld.DateBirthday)
	if result.Error != nil {
		return check(op, errors.New("actor doesn't exist"))
	}
	if actor.ValidateGender() {
		resultUpdate := p.Db.Model(model.Actor{}).Where("name = ? AND date_birthday = ?", actorOld.Name, actorOld.DateBirthday).Updates(model.Actor{Gender: actor.Gender})
		if resultUpdate.Error != nil {
			return check(op, errors.New("failed to update actor"))
		}
	}
	if actor.ValidateName() {
		resultUpdate := p.Db.Model(model.Actor{}).Where("name = ? AND date_birthday = ?", actorOld.Name, actorOld.DateBirthday).Updates(model.Actor{Name: actor.Name})
		if resultUpdate.Error != nil {
			return check(op, errors.New("failed to update actor"))
		}
	}
	if actor.ValidateBirthday() {
		resultUpdate := p.Db.Model(model.Actor{}).Where("name = ? AND date_birthday = ?", actor.Name, actorOld.DateBirthday).Updates(model.Actor{DateBirthday: actor.DateBirthday})
		if resultUpdate.Error != nil {
			return check(op, errors.New("failed to update actor"))
		}
	}
	return nil
}

func (p *Postgres) SelectAllActors() ([]models.ActorFilm, error) {
	const op = "repository.postgre.SelectAllActors"
	actors := make(map[int]models.ActorFilm)
	films := make(map[int][]string)
	var result []models.ActorFilm
	rows, err := p.Db.Table("actor").Select("actor.id, name, title, gender, date_birthday").Joins("left join actors_films on id = actors_films.actor_id").Joins("left join films on films.id = actors_films.film_id").Rows()
	if err != nil {
		return nil, check(op, errors.New("cannot find results"))
	}
	defer rows.Close()
	for rows.Next() {
		var id *int
		var name *string
		var title *string
		var gender *string
		var date *strfmt.Date
		rows.Scan(&id, &name, &title, &gender, &date)
		if id != nil {
			_, ok := actors[*id]
			if !ok {
				actors[*id] = models.ActorFilm{Name: name, Sex: gender, DateOfBirthday: date}
			}
			if title != nil {
				films[*id] = append(films[*id], *title)
			}
		}
	}
	i := 0
	for key := range actors {
		result = append(result, actors[key])
		result[i].Films = films[key]
		i++
	}
	return result, nil
}

func (p *Postgres) InsertFilm(film model.Film, actors []model.Actor) error {
	const op = "repository.postgre.InsertFilm"
	var films model.Film
	p.Db.Where("title = ? AND date_premiere = ?", film.Title, film.DatePremiere).First(&films)
	if films.ValidateTitle() && films.ValidateDatePremiere() {
		return check(op, errors.New("this film already exist"))
	}
	result := p.Db.Create(&film)
	if result.Error != nil {
		return check(op, errors.New("faile to create film"))
	}
	for _, val := range actors {
		p.Db.Create(&val)
		p.Db.Table("actors_films").Create(&model.Actors_films{val.Id, film.Id})
	}
	return nil

}

func (p *Postgres) DeleteFilm(title string, premiere int) error {
	const op = "repository.postgre.DeleteFilm"
	var record model.Film
	result := p.Db.Model(model.Film{}).Where("title = ? AND date_premiere = ?", title, premiere).Delete(record)
	var err error = nil
	if result.Error != nil {
		err = errors.New("failed to delete film")
	}
	var value model.Actors_films
	p.Db.Table("actors_films").Where("film_id = ?", record.Id).Delete(value)
	return check(op, err)
}

func (p *Postgres) UpdateFilm(filmOld, film model.Film) error {
	const op = "repository.postgre.UpdateFilm"
	result := p.Db.Model(model.Film{}).Where("title = ? ", filmOld.Title)
	if result.Error != nil {
		return check(op, errors.New("film doesn't exist"))
	}
	if film.ValidateTitle() {
		resultUpdate := p.Db.Model(model.Film{}).Where("title = ? ", filmOld.Title).Updates(model.Film{Title: film.Title})
		if resultUpdate.Error != nil {
			return check(op, errors.New("failed to update film"))
		}
	}
	if film.ValidateDescription() {
		resultUpdate := p.Db.Model(model.Film{}).Where("title = ? ", filmOld.Title).Updates(model.Film{Description: film.Description})
		if resultUpdate.Error != nil {
			return check(op, errors.New("failed to update film"))
		}
	}
	if film.ValidateDatePremiere() {
		resultUpdate := p.Db.Model(model.Film{}).Where("title = ? ", filmOld.Title).Updates(model.Film{DatePremiere: film.DatePremiere})
		if resultUpdate.Error != nil {
			return check(op, errors.New("failed to update film"))
		}
	}
	if film.ValidateRating() {
		resultUpdate := p.Db.Model(model.Film{}).Where("title = ? ", filmOld.Title).Updates(model.Film{Rating: film.Rating})
		if resultUpdate.Error != nil {
			return check(op, errors.New("failed to update film"))
		}
	}
	return nil
}

func (p *Postgres) SelectFilmsPattern(filmName string, actorName string) ([]model.Film, error) {
	const op = "repository.postgre.SelectFilmsPattern"
	querys := make([]model.Film, 0)
	rows, err := p.Db.Table("actor").Where("title LIKE ? AND name LIKE ?", "%"+filmName+"%", "%"+actorName+"%").Select("title, description, date_premiere, rating").Joins("join actors_films on id = actors_films.actor_id").Joins("join films on films.id = actors_films.film_id").Rows()
	if err != nil {
		return nil, check(op, errors.New("fail to find results"))
	}
	defer rows.Close()
	for rows.Next() {
		var film model.Film
		rows.Scan(&film.Title, &film.Description, &film.DatePremiere, &film.Rating)
		querys = append(querys, film)
	}
	return querys, nil
}

func (p *Postgres) SelectSortedFilms(orderBy string) ([]model.Film, error) {
	const op = "repository.postgre.SelectSortedFilms"
	querys := make([]model.Film, 0)
	rows := p.Db.Table("films").Order(orderBy + " Desc").Find(&querys)
	if rows.Error != nil {
		return nil, check(op, errors.New("failed to find films"))
	}
	return querys, nil
}

func check(op string, err error) error {
	if err != nil {
		err = fmt.Errorf("%s: %w", op, err)
	}
	return err
}
