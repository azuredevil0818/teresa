package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
	"github.com/luizalabs/teresa/api/models"
)

// GetDeploymentsHandlerFunc turns a function with the right signature into a get deployments handler
type GetDeploymentsHandlerFunc func(GetDeploymentsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDeploymentsHandlerFunc) Handle(params GetDeploymentsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetDeploymentsHandler interface for that can handle valid get deployments params
type GetDeploymentsHandler interface {
	Handle(GetDeploymentsParams, interface{}) middleware.Responder
}

// NewGetDeployments creates a new http.Handler for the get deployments operation
func NewGetDeployments(ctx *middleware.Context, handler GetDeploymentsHandler) *GetDeployments {
	return &GetDeployments{Context: ctx, Handler: handler}
}

/*GetDeployments swagger:route GET /teams/{team_id}/apps/{app_id}/deployments deployments getDeployments

Get app deployments

Get deployments an app may have

*/
type GetDeployments struct {
	Context *middleware.Context
	Handler GetDeploymentsHandler
}

func (o *GetDeployments) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetDeploymentsParams()

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

/*GetDeploymentsOKBodyBody get deployments o k body body

swagger:model GetDeploymentsOKBodyBody
*/
type GetDeploymentsOKBodyBody struct {
	models.Pagination

	/* items

	Required: true
	*/
	Items []*models.Deployment `json:"items"`
}

// Validate validates this get deployments o k body body
func (o *GetDeploymentsOKBodyBody) Validate(formats strfmt.Registry) error {
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

func (o *GetDeploymentsOKBodyBody) validateItems(formats strfmt.Registry) error {

	if err := validate.Required("getDeploymentsOK"+"."+"items", "body", o.Items); err != nil {
		return err
	}

	return nil
}
