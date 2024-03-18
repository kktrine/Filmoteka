// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"filmoteka_server/internal/auth"
	"filmoteka_server/internal/config"
	"filmoteka_server/internal/log"
	"filmoteka_server/internal/model"
	"filmoteka_server/internal/repository/postgre"
	"filmoteka_server/models"
	"filmoteka_server/restapi/operations"
	"filmoteka_server/restapi/operations/actor"
	"filmoteka_server/restapi/operations/film"
	"filmoteka_server/restapi/operations/user"
	"log/slog"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	//"log"
	"net/http"
)

func configureFlags(api *operations.FilmotekaAPI) {
}

func configureAPI(api *operations.FilmotekaAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError
	cfg := config.MustLoad()
	logger := log.SetupLogger(cfg.LogPath)
	api.Logger = logger.Info
	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "admin-key" header is set
	api.IsAdminAuth = func(token string) (*models.Principal, error) {
		return auth.CheckAdminToken(token)
	}
	// Applies when the "user-key" header is set

	api.IsUserAuth = func(token string) (*models.Principal, error) {
		return auth.CheckUserToken(token)
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	storage := postgre.NewPostgresRepository(cfg.DbConfig)

	api.ActorAddActorHandler = actor.AddActorHandlerFunc(func(params actor.AddActorParams, principal *models.Principal) middleware.Responder {
		actorToInsert := model.Actor{
			Name:         *params.Body.Name,
			Gender:       *params.Body.Sex,
			DateBirthday: time.Time(*(params.Body.DateOfBirthday)),
		}
		err := storage.InsertActor(actorToInsert)
		if err != nil {
			return &actor.AddActorBadRequest{
				Payload: err.Error(),
			}
		}
		return &actor.AddActorOK{}
	})

	api.FilmAddFilmHandler = film.AddFilmHandlerFunc(func(params film.AddFilmParams, principal *models.Principal) middleware.Responder {
		filmToInsert := model.Film{
			Title:        *params.Body.Name,
			Description:  *params.Body.Description,
			DatePremiere: int(*params.Body.Year),
			Rating:       *params.Body.Rate,
		}
		var actorsToInsert []model.Actor
		for _, value := range params.Body.Actors {
			curr := model.Actor{Name: *value.Name, Gender: *value.Sex, DateBirthday: time.Time(*value.DateOfBirthday)}
			actorsToInsert = append(actorsToInsert, curr)
		}
		err := storage.InsertFilm(filmToInsert, actorsToInsert)
		if err != nil {
			return &film.AddFilmBadRequest{Payload: err.Error()}
		}
		return &film.AddFilmOK{}
	})

	api.ActorDeleteActorHandler = actor.DeleteActorHandlerFunc(func(params actor.DeleteActorParams, principal *models.Principal) middleware.Responder {
		actorToDelete := model.Actor{
			Name:         params.Name,
			DateBirthday: time.Time(params.BDate),
		}
		err := storage.DeleteActor(actorToDelete)
		if err != nil {
			return &actor.DeleteActorBadRequest{
				Payload: err.Error(),
			}
		}
		return &actor.DeleteActorNoContent{}
	})

	api.FilmDeleteFilmHandler = film.DeleteFilmHandlerFunc(func(params film.DeleteFilmParams, principal *models.Principal) middleware.Responder {
		err := storage.DeleteFilm(params.Name, int(params.Year))
		if err != nil {
			return &film.DeleteFilmBadRequest{Payload: err.Error()}
		}
		return &film.DeleteFilmNoContent{}
	})

	api.FilmFindFilmHandler = film.FindFilmHandlerFunc(func(params film.FindFilmParams, principal *models.Principal) middleware.Responder {
		data, err := storage.SelectFilmsPattern(params.Actor, params.Desc)
		//fmt.Println(data)
		if err != nil {
			return &film.FindFilmBadRequest{Payload: err.Error()}
		}
		var res []*film.FindFilmOKBodyItems0
		for _, curr := range data {
			year := int64(curr.DatePremiere)
			res = append(res, &film.FindFilmOKBodyItems0{
				Name:        &curr.Title,
				Year:        &year,
				Description: &curr.Description,
			})
		}
		return &film.FindFilmOK{Payload: res}
	})

	api.ActorGetAllActorsHandler = actor.GetAllActorsHandlerFunc(func(params actor.GetAllActorsParams, principal *models.Principal) middleware.Responder {
		data, err := storage.SelectAllActors()
		if err != nil {
			return &actor.GetAllActorsBadRequest{
				Payload: err.Error(),
			}
		}
		var res []*actor.GetAllActorsOKBodyItems0
		for _, record := range data {
			curr := actor.GetAllActorsOKBodyItems0{Name: *record.Name, DateOfBirthday: *record.DateOfBirthday}
			for _, val := range record.Films {
				curr.Films = append(curr.Films, val)
			}
			res = append(res, &curr)
		}
		return &actor.GetAllActorsOK{
			Payload: res,
		}
	})

	api.FilmGetAllFilmsSortedHandler = film.GetAllFilmsSortedHandlerFunc(func(params film.GetAllFilmsSortedParams, principal *models.Principal) middleware.Responder {
		data, err := storage.SelectSortedFilms("title")
		if err != nil {
			return &film.GetAllFilmsSortedBadRequest{
				Payload: err.Error(),
			}
		}
		var res []*film.GetAllFilmsSortedOKBodyItems0
		for _, curr := range data {
			year := int64(curr.DatePremiere)
			res = append(res, &film.GetAllFilmsSortedOKBodyItems0{
				Description: &curr.Description,
				Name:        &curr.Title,
				Rating:      &curr.Rating,
				Year:        &year,
			})
		}
		return &film.GetAllFilmsSortedOK{Payload: res}
	})

	api.ActorUpdateActorHandler = actor.UpdateActorHandlerFunc(func(params actor.UpdateActorParams, principal *models.Principal) middleware.Responder {
		actorToUpdate := model.Actor{
			Name:         params.Name,
			DateBirthday: time.Time(params.BDate),
		}
		newValue := model.Actor{}
		if sex := params.Body.Sex; sex != "" {
			newValue.Gender = sex
		}
		if name := params.Body.Name; name != "" {
			newValue.Name = name
		}
		var tmp string
		zeroDate, _ := time.Parse("0001-01-01", tmp)
		if year := params.Body.DateOfBirthday; time.Time(year) != zeroDate {
			newValue.DateBirthday = time.Time(year)
		}
		//fmt.Println(newValue)
		err := storage.UpdateActor(actorToUpdate, newValue)
		if err != nil {
			return &actor.UpdateActorBadRequest{Payload: err.Error()}
		}
		return &actor.UpdateActorOK{}
	})

	api.FilmUpdateFilmHandler = film.UpdateFilmHandlerFunc(func(params film.UpdateFilmParams, principal *models.Principal) middleware.Responder {
		filmToUpdate := model.Film{
			Title:        params.Name,
			DatePremiere: int(params.Year),
		}
		newValue := model.Film{}
		if name := params.Body.Name; name != "" {
			newValue.Title = name
		}
		if year := params.Body.Year; year != 0 {
			newValue.DatePremiere = int(year)
		}
		if desc := params.Body.Description; desc != "" {
			newValue.Description = desc
		}
		if rate := params.Body.Rate; rate != nil {
			newValue.Rating = *rate
		}
		err := storage.UpdateFilm(filmToUpdate, newValue)
		if err != nil {
			return &film.UpdateFilmBadRequest{Payload: err.Error()}
		}

		return &film.UpdateFilmOK{}
	})

	api.UserUserLoginHandler = user.UserLoginHandlerFunc(func(params user.UserLoginParams) middleware.Responder {
		token, err := auth.Login(params.Username, params.Password, storage)
		if err != nil {
			return &user.UserLoginBadRequest{}
		}
		return &user.UserLoginOK{Payload: token}
	})

	api.UserUserLogoutHandler = user.UserLogoutHandlerFunc(func(params user.UserLogoutParams, principal *models.Principal) middleware.Responder {
		auth.Logout(params.Token)
		return &user.UserLogoutOK{}
	})

	api.PreServerShutdown = func() { storage.Stop() }

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(logger, api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(logger *slog.Logger, handler http.Handler) http.Handler {
	return addLogging(logger)(handler)
}

func addLogging(log *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		log := log.With(
			slog.String("component", "middleware/logger"),
		)
		log.Info("logger middleware enabled")
		fn := func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()
			ww := &responseWriterWithStatus{w, http.StatusOK}
			defer func() {
				status := ww.Status()
				elapsed := time.Since(startTime)
				log.Info("Request status", slog.String("Path", r.URL.Path), slog.String("Method", r.Method), slog.Int("status_code", status), slog.Duration("elapsed_time", elapsed))
			}()
			next.ServeHTTP(ww, r)
		}

		return http.HandlerFunc(fn)
	}
}

type responseWriterWithStatus struct {
	http.ResponseWriter
	status int
}

func (r *responseWriterWithStatus) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

func (r *responseWriterWithStatus) Status() int {
	return r.status
}
