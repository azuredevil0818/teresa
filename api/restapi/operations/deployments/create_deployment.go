package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateDeploymentHandlerFunc turns a function with the right signature into a create deployment handler
type CreateDeploymentHandlerFunc func(CreateDeploymentParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateDeploymentHandlerFunc) Handle(params CreateDeploymentParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateDeploymentHandler interface for that can handle valid create deployment params
type CreateDeploymentHandler interface {
	Handle(CreateDeploymentParams, interface{}) middleware.Responder
}

// NewCreateDeployment creates a new http.Handler for the create deployment operation
func NewCreateDeployment(ctx *middleware.Context, handler CreateDeploymentHandler) *CreateDeployment {
	return &CreateDeployment{Context: ctx, Handler: handler}
}

/*CreateDeployment swagger:route POST /teams/{team_id}/apps/{app_id}/deployments deployments createDeployment

Create a new deploy

Create a deploy for an app

*/
type CreateDeployment struct {
	Context *middleware.Context
	Handler CreateDeploymentHandler
}

func (o *CreateDeployment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewCreateDeploymentParams()

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
