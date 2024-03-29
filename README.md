# Filmoteka_server

Спецификация swagger 2.0. Api-first подход

Версия go - 1.22


# Алгоритм установки и запуска проекта:
В файлы `.yaml` директории config, а также в файлы `docker-compose.yml`, `wait_for_postgres.sh`  необходимо внести имя пользователя и пароль для пользователя базы данных.

Запуск:
- Через Docker `docker-compose up --build filmoteka-server`
- Локально `make run_service`

Запуск тестов: из директории `iternal/repository/postgre`: `go test`



# Конфигурация базы данных

Для работы с базой данных используется пакет gorm

В базе данных присутствует 4 таблицы:
actor - хранит актеров
films - хранит фильмы
actors_films - связывает фильмы и актеров
users - хранит пользователей(Сервер кроме Login не предоставляет операции с данной таблицей). В ней прописаны два пользователя с ролями isUser, isAdmin

# Первый запуск

При первом запуске сервера будут автоматически созданы все таблицы из базы данных.
Данные заносятся через incert.sql скрипт


# Основные операции
Проект filmoteka позволяет пользователям хранить информацию о фильмах и актерах, а так же получать эту информацию по запросу. В качестве хранилища данных используется БД PostgreSQL.

Проект запускается в докер контейнерах: предоставлен dockerfile, docker-compose для сборки проекта в докере. Так же предоставлен скрипт, который ожидает полного включения базы данных в контейнере перед запуском контейнера с приложением.

Имеется два типа пользователь. Сначала необходимо аутентифицироваться по пути Filmoteka/user/login с параметрами username и password. 
В базе данных заранее подготовлено 2 пользователя: \
**isAdmin** username: **user1** password: **password1** \
**isUser** username: **user2** password: **password2** 

Обычный пользователь (_isUser_) может совершать GET запросы.
Пользователь с правами администратора (_isAdmin_) может совершать любые доступные запросы.

### Запросы на получение данных:
- Получать из БД список актеров и список фильмов с их участием `GET /actors`
- Получать информацию о фильмах из БД по фрагменту названия фильма и фрагменту имени актера `GET /films/find`
- Получать список фильмов из БД с возможностью сортировки по названию, рейтингу, дате выпуска `GET /films`

### Запросы на аутентификацию
- Получить токен пользователя или админа (зависит от логина) `GET /user/login`\
Далее в каждом запросе необходимо посылать `user-key` / `admin-key`
- Выйти из аккаунта (удалить токен из системы) `GET /user/logout`

### Запросы модифицирующие данные

- Добавлять информацию в БД об актерах (имя, пол, дата рождения) `POST /actors`
- Изменять информацию об актерах (любое из полей или несколько полей) `PUT /actors`
- Удалять информацию об актере по его имени и дате рождения `DELETE /actors`
- Добавлять фильмы и информацию о фильмах в БД `POST /films`
- Изменять информацию о фильмах (любое из полей или несколько полей) `PUT /films`







