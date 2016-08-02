package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteUserParams creates a new DeleteUserParams object
// with the default values initialized.
func NewDeleteUserParams() *DeleteUserParams {
	var ()
	return &DeleteUserParams{}
}

/*DeleteUserParams contains all the parameters to send to the API endpoint
for the delete user operation typically these are written to a http.Request
*/
type DeleteUserParams struct {

	/*UserID
	  User ID

	*/
	UserID int64
}

// WithUserID adds the userId to the delete user params
func (o *DeleteUserParams) WithUserID(UserID int64) *DeleteUserParams {
	o.UserID = UserID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
