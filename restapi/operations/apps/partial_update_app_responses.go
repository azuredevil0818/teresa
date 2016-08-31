package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/luizalabs/teresa-api/models"
)

/*PartialUpdateAppOK App

swagger:response partialUpdateAppOK
*/
type PartialUpdateAppOK struct {

	// In: body
	Payload *models.App `json:"body,omitempty"`
}

// NewPartialUpdateAppOK creates PartialUpdateAppOK with default headers values
func NewPartialUpdateAppOK() *PartialUpdateAppOK {
	return &PartialUpdateAppOK{}
}

// WithPayload adds the payload to the partial update app o k response
func (o *PartialUpdateAppOK) WithPayload(payload *models.App) *PartialUpdateAppOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the partial update app o k response
func (o *PartialUpdateAppOK) SetPayload(payload *models.App) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PartialUpdateAppOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PartialUpdateAppUnauthorized User not authorized

swagger:response partialUpdateAppUnauthorized
*/
type PartialUpdateAppUnauthorized struct {

	// In: body
	Payload *models.Unauthorized `json:"body,omitempty"`
}

// NewPartialUpdateAppUnauthorized creates PartialUpdateAppUnauthorized with default headers values
func NewPartialUpdateAppUnauthorized() *PartialUpdateAppUnauthorized {
	return &PartialUpdateAppUnauthorized{}
}

// WithPayload adds the payload to the partial update app unauthorized response
func (o *PartialUpdateAppUnauthorized) WithPayload(payload *models.Unauthorized) *PartialUpdateAppUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the partial update app unauthorized response
func (o *PartialUpdateAppUnauthorized) SetPayload(payload *models.Unauthorized) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PartialUpdateAppUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PartialUpdateAppForbidden User does not have the credentials to access this resource


swagger:response partialUpdateAppForbidden
*/
type PartialUpdateAppForbidden struct {

	// In: body
	Payload *models.Unauthorized `json:"body,omitempty"`
}

// NewPartialUpdateAppForbidden creates PartialUpdateAppForbidden with default headers values
func NewPartialUpdateAppForbidden() *PartialUpdateAppForbidden {
	return &PartialUpdateAppForbidden{}
}

// WithPayload adds the payload to the partial update app forbidden response
func (o *PartialUpdateAppForbidden) WithPayload(payload *models.Unauthorized) *PartialUpdateAppForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the partial update app forbidden response
func (o *PartialUpdateAppForbidden) SetPayload(payload *models.Unauthorized) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PartialUpdateAppForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PartialUpdateAppDefault Error

swagger:response partialUpdateAppDefault
*/
type PartialUpdateAppDefault struct {
	_statusCode int

	// In: body
	Payload *models.GenericError `json:"body,omitempty"`
}

// NewPartialUpdateAppDefault creates PartialUpdateAppDefault with default headers values
func NewPartialUpdateAppDefault(code int) *PartialUpdateAppDefault {
	if code <= 0 {
		code = 500
	}

	return &PartialUpdateAppDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the partial update app default response
func (o *PartialUpdateAppDefault) WithStatusCode(code int) *PartialUpdateAppDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the partial update app default response
func (o *PartialUpdateAppDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the partial update app default response
func (o *PartialUpdateAppDefault) WithPayload(payload *models.GenericError) *PartialUpdateAppDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the partial update app default response
func (o *PartialUpdateAppDefault) SetPayload(payload *models.GenericError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PartialUpdateAppDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
