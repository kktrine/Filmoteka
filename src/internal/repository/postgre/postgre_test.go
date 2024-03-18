package postgre

import (
	_ "database/sql"
	"errors"
	"filmoteka_server/internal/model"

	"log"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewMockDB() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if db == nil {
		log.Fatal("mock db is null")
	}
	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	dialector := postgres.New(postgres.Config{
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})

	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening gorm database", err)
	}

	return gormDB, mock
}

func TestLoginFirst(t *testing.T) {
	db, mock := NewMockDB()

	rows := sqlmock.NewRows([]string{"user_id", "name", "password", "role"}).
		AddRow(1, "alexey", "12345", "admin")

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"users\" WHERE name = $1 LIMIT $2")).WillReturnRows(rows)

	var users model.User
	dataBase := &Postgres{Db: db}
	val, err := dataBase.Login("alexey", "12345")
	if err != nil {
		t.Fatalf("Error in finding users: %v", err)
	}

	if val != "admin" {
		t.Fatalf("Unexpected user data retrieved: %v", users)
	}
}

func TestLoginSecond(t *testing.T) {
	db, mock := NewMockDB()

	rows := sqlmock.NewRows([]string{"user_id", "name", "password", "role"}).
		AddRow(1, "alexey", "12345", "admin")

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"users\" WHERE name = $1 LIMIT $2")).WillReturnRows(rows)

	dataBase := &Postgres{Db: db}
	_, err := dataBase.Login("alexey", "123456")
	assert.EqualError(t, err, "repository.postgre.Login: incorrect pasword or login")
}

func TestInsertActorsFirst(t *testing.T) {
	db, mock := NewMockDB()

	rows := sqlmock.NewRows([]string{"id", "name", "gender", "date_birthday"}).
		AddRow(1, "alexey", "male", time.Date(2022, time.March, 15, 0, 0, 0, 0, time.UTC))

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"actor\" WHERE name = $1 AND date_birthday = $2 ORDER BY \"actor\".\"id\" LIMIT $3")).WillReturnRows(rows)

	dataBase := &Postgres{Db: db}
	actor := model.Actor{
		Id:           1,
		Name:         "alexey",
		Gender:       "male",
		DateBirthday: time.Date(2022, time.March, 15, 0, 0, 0, 0, time.UTC),
	}
	err := dataBase.InsertActor(actor)
	assert.EqualError(t, err, "repository.postgre.InsertActor: this actor already exist")
}

func TestInsertActorsSecond(t *testing.T) {
	db, mock := NewMockDB()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"actor\" WHERE name = $1 AND date_birthday = $2 ORDER BY \"actor\".\"id\" LIMIT $3")).WillReturnError(errors.New("error"))
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO \"actors\"")).WillReturnResult(sqlmock.NewResult(1, 1))
	dataBase := &Postgres{Db: db}
	actor := model.Actor{
		Id:           2,
		Name:         "leha",
		Gender:       "female",
		DateBirthday: time.Date(2023, time.March, 15, 0, 0, 0, 0, time.UTC),
	}
	err := dataBase.InsertActor(actor)
	assert.EqualError(t, err, "repository.postgre.InsertActor: faile to create actor")
}

func TestInsertActorThird(t *testing.T) {
	db, mock := NewMockDB()

	actor := model.Actor{
		Id:           2,
		Name:         "leha",
		Gender:       "female",
		DateBirthday: time.Date(2023, time.March, 15, 0, 0, 0, 0, time.UTC),
	}
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO \"actor\" (\"name\",\"gender\",\"date_birthday\",\"id\") VALUES ($1,$2,$3,$4) RETURNING \"id\"")).
		WithArgs(actor.Name, actor.Gender, actor.DateBirthday, actor.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "gender", "date_birthday"}))
	mock.ExpectCommit()
	dataBase := &Postgres{Db: db}
	if err := dataBase.InsertActor(actor); err != nil {
		t.Fatalf("Failed to insert user: %v", err)
	}
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteActor(t *testing.T) {
	db, mock := NewMockDB()
	actor := model.Actor{
		Id:           1,
		Name:         "alexey",
		Gender:       "male",
		DateBirthday: time.Date(2023, time.March, 15, 0, 0, 0, 0, time.UTC),
	}
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM \"actor\" WHERE name = $1 AND date_birthday = $2")).
		WithArgs(actor.Name, actor.DateBirthday).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	dataBase := &Postgres{Db: db}
	err := dataBase.DeleteActor(actor)
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSelectAllActors(t *testing.T) {
	db, mock := NewMockDB()

	rows := sqlmock.NewRows([]string{"id", "name", "gender", "date_birthday"}).
		AddRow(1, "alexey", "male", time.Date(2022, time.March, 15, 0, 0, 0, 0, time.UTC))

	mock.ExpectQuery("SELECT actor.id, name, title, gender, date_birthday FROM \"actor\" left join actors_films on id = actors_films.actor_id left join films on films.id = actors_films.film_id").WillReturnRows(rows)

	dataBase := &Postgres{Db: db}
	_, err := dataBase.SelectAllActors()
	if err != nil {
		t.Fatalf("Error in finding actors: %v", err)
	}
}

func TestDeleteFilm(t *testing.T) {
	db, mock := NewMockDB()
	film := model.Film{
		Id:           2,
		Title:        "Best Film",
		Description:  "very intresting film",
		DatePremiere: 2026,
		Rating:       10,
	}
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM \"films\" WHERE title = $1 AND date_premiere = $2")).
		WithArgs(film.Title, film.DatePremiere).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	dataBase := &Postgres{Db: db}
	err := dataBase.DeleteFilm(film.Title, film.DatePremiere)
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, mock.ExpectationsWereMet())
}
