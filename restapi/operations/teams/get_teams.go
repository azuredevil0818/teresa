package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
	"github.com/luizalabs/tapi/models"
)

// GetTeamsHandlerFunc turns a function with the right signature into a get teams handler
type GetTeamsHandlerFunc func(GetTeamsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTeamsHandlerFunc) Handle(params GetTeamsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetTeamsHandler interface for that can handle valid get teams params
type GetTeamsHandler interface {
	Handle(GetTeamsParams, interface{}) middleware.Responder
}

// NewGetTeams creates a new http.Handler for the get teams operation
func NewGetTeams(ctx *middleware.Context, handler GetTeamsHandler) *GetTeams {
	return &GetTeams{Context: ctx, Handler: handler}
}

/*GetTeams swagger:route GET /teams teams getTeams

Get teams

Find and filter teams

*/
type GetTeams struct {
	Context *middleware.Context
	Handler GetTeamsHandler
}

func (o *GetTeams) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetTeamsParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

/*GetTeamsOKBodyBody get teams o k body body

swagger:model GetTeamsOKBodyBody
*/
type GetTeamsOKBodyBody struct {
	models.Pagination

	/* items

	Required: true
	*/
	Items []*models.Team `json:"items"`
}

// Validate validates this get teams o k body body
func (o *GetTeamsOKBodyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.Pagination.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTeamsOKBodyBody) validateItems(formats strfmt.Registry) error {

	if err := validate.Required("getTeamsOK"+"."+"items", "body", o.Items); err != nil {
		return err
	}

	for i := 0; i < len(o.Items); i++ {

		if swag.IsZero(o.Items[i]) { // not required
			continue
		}

		if o.Items[i] != nil {

			if err := o.Items[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
