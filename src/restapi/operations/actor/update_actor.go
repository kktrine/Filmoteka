// Code generated by go-swagger; DO NOT EDIT.

package actor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"filmoteka_server/models"
)

// UpdateActorHandlerFunc turns a function with the right signature into a update actor handler
type UpdateActorHandlerFunc func(UpdateActorParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateActorHandlerFunc) Handle(params UpdateActorParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// UpdateActorHandler interface for that can handle valid update actor params
type UpdateActorHandler interface {
	Handle(UpdateActorParams, *models.Principal) middleware.Responder
}

// NewUpdateActor creates a new http.Handler for the update actor operation
func NewUpdateActor(ctx *middleware.Context, handler UpdateActorHandler) *UpdateActor {
	return &UpdateActor{Context: ctx, Handler: handler}
}

/*
	UpdateActor swagger:route PUT /actors Actor updateActor

Изменить информацию об актере
*/
type UpdateActor struct {
	Context *middleware.Context
	Handler UpdateActorHandler
}

func (o *UpdateActor) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateActorParams()
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

// UpdateActorBody update actor body
//
// swagger:model UpdateActorBody
type UpdateActorBody struct {

	// date of birthday
	// Format: date
	DateOfBirthday strfmt.Date `json:"date_of_birthday,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// sex
	// Enum: ["M","F"]
	Sex string `json:"sex,omitempty"`
}

// Validate validates this update actor body
func (o *UpdateActorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDateOfBirthday(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSex(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateActorBody) validateDateOfBirthday(formats strfmt.Registry) error {
	if swag.IsZero(o.DateOfBirthday) { // not required
		return nil
	}

	if err := validate.FormatOf("body"+"."+"date_of_birthday", "body", "date", o.DateOfBirthday.String(), formats); err != nil {
		return err
	}

	return nil
}

var updateActorBodyTypeSexPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["M","F"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateActorBodyTypeSexPropEnum = append(updateActorBodyTypeSexPropEnum, v)
	}
}

const (

	// UpdateActorBodySexM captures enum value "M"
	UpdateActorBodySexM string = "M"

	// UpdateActorBodySexF captures enum value "F"
	UpdateActorBodySexF string = "F"
)

// prop value enum
func (o *UpdateActorBody) validateSexEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateActorBodyTypeSexPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateActorBody) validateSex(formats strfmt.Registry) error {
	if swag.IsZero(o.Sex) { // not required
		return nil
	}

	// value enum
	if err := o.validateSexEnum("body"+"."+"sex", "body", o.Sex); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update actor body based on context it is used
func (o *UpdateActorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateActorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateActorBody) UnmarshalBinary(b []byte) error {
	var res UpdateActorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}