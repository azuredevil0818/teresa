package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUpdateUserParams creates a new UpdateUserParams object
// with the default values initialized.
func NewUpdateUserParams() *UpdateUserParams {
	var ()
	return &UpdateUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUserParamsWithTimeout creates a new UpdateUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateUserParamsWithTimeout(timeout time.Duration) *UpdateUserParams {
	var ()
	return &UpdateUserParams{

		timeout: timeout,
	}
}

/*UpdateUserParams contains all the parameters to send to the API endpoint
for the update user operation typically these are written to a http.Request
*/
type UpdateUserParams struct {

	/*UserID
	  User ID

	*/
	UserID int64

	timeout time.Duration
}

// WithUserID adds the userID to the update user params
func (o *UpdateUserParams) WithUserID(userID int64) *UpdateUserParams {
	o.UserID = userID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param user_id
	if err := r.SetPathParam("user_id", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
