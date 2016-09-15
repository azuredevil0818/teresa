package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*App app

swagger:model App
*/
type App struct {
	AppIn

	/* list of the external addresses of the app
	 */
	AddressList []string `json:"addressList,omitempty"`

	/* creator

	Required: true
	*/
	Creator *User `json:"creator"`

	/* deployment list
	 */
	DeploymentList []*Deployment `json:"deploymentList,omitempty"`

	/* env vars
	 */
	EnvVars []*EnvVar `json:"envVars,omitempty"`
}

// Validate validates this app
func (m *App) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.AppIn.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAddressList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeploymentList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvVars(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *App) validateAddressList(formats strfmt.Registry) error {

	return nil
}

func (m *App) validateCreator(formats strfmt.Registry) error {

	if err := validate.Required("creator", "body", m.Creator); err != nil {
		return err
	}

	if m.Creator != nil {

		if err := m.Creator.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *App) validateDeploymentList(formats strfmt.Registry) error {

	for i := 0; i < len(m.DeploymentList); i++ {

		if swag.IsZero(m.DeploymentList[i]) { // not required
			continue
		}

		if m.DeploymentList[i] != nil {

			if err := m.DeploymentList[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *App) validateEnvVars(formats strfmt.Registry) error {

	for i := 0; i < len(m.EnvVars); i++ {

		if swag.IsZero(m.EnvVars[i]) { // not required
			continue
		}

		if m.EnvVars[i] != nil {

			if err := m.EnvVars[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
