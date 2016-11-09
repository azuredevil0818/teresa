package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/luizalabs/teresa-api/models"
)

// NewPartialUpdateAppParams creates a new PartialUpdateAppParams object
// with the default values initialized.
func NewPartialUpdateAppParams() PartialUpdateAppParams {
	var ()
	return PartialUpdateAppParams{}
}

// PartialUpdateAppParams contains all the bound params for the partial update app operation
// typically these are obtained from a http.Request
//
// swagger:parameters partialUpdateApp
type PartialUpdateAppParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*App ID
	  Required: true
	  In: path
	*/
	AppID int64
	/*
	  Required: true
	  In: body
	*/
	Body []*models.PatchAppRequest
	/*Team ID
	  Required: true
	  In: path
	*/
	TeamID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *PartialUpdateAppParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rAppID, rhkAppID, _ := route.Params.GetOK("app_id")
	if err := o.bindAppID(rAppID, rhkAppID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body []*models.PatchAppRequest
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body"))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}

		} else {
			for _, io := range o.Body {
				if err := io.Validate(route.Formats); err != nil {
					res = append(res, err)
					break
				}
			}

			if len(res) == 0 {
				o.Body = body
			}
		}

	} else {
		res = append(res, errors.Required("body", "body"))
	}

	rTeamID, rhkTeamID, _ := route.Params.GetOK("team_id")
	if err := o.bindTeamID(rTeamID, rhkTeamID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PartialUpdateAppParams) bindAppID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("app_id", "path", "int64", raw)
	}
	o.AppID = value

	return nil
}

func (o *PartialUpdateAppParams) bindTeamID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("team_id", "path", "int64", raw)
	}
	o.TeamID = value

	return nil
}
