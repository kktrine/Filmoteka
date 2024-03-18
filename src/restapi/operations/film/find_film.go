// Code generated by go-swagger; DO NOT EDIT.

package film

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"filmoteka_server/models"
)

// FindFilmHandlerFunc turns a function with the right signature into a find film handler
type FindFilmHandlerFunc func(FindFilmParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn FindFilmHandlerFunc) Handle(params FindFilmParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// FindFilmHandler interface for that can handle valid find film params
type FindFilmHandler interface {
	Handle(FindFilmParams, *models.Principal) middleware.Responder
}

// NewFindFilm creates a new http.Handler for the find film operation
func NewFindFilm(ctx *middleware.Context, handler FindFilmHandler) *FindFilm {
	return &FindFilm{Context: ctx, Handler: handler}
}

/*
	FindFilm swagger:route GET /films/find Film findFilm

Поиск фильма по фрагменту описания или имени актера
*/
type FindFilm struct {
	Context *middleware.Context
	Handler FindFilmHandler
}

func (o *FindFilm) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewFindFilmParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// FindFilmOKBodyItems0 find film o k body items0
//
// swagger:model FindFilmOKBodyItems0
type FindFilmOKBodyItems0 struct {

	// description
	// Required: true
	// Max Length: 1000
	Description *string `json:"description"`

	// name
	// Required: true
	Name *string `json:"name"`

	// year
	// Required: true
	Year *int64 `json:"year"`
}

// Validate validates this find film o k body items0
func (o *FindFilmOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateYear(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FindFilmOKBodyItems0) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", o.Description); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", *o.Description, 1000); err != nil {
		return err
	}

	return nil
}

func (o *FindFilmOKBodyItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *FindFilmOKBodyItems0) validateYear(formats strfmt.Registry) error {

	if err := validate.Required("year", "body", o.Year); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this find film o k body items0 based on context it is used
func (o *FindFilmOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *FindFilmOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *FindFilmOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res FindFilmOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
