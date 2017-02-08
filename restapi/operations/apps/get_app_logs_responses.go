package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/luizalabs/teresa-api/models"
)

/*GetAppLogsOK App logs

swagger:response getAppLogsOK
*/
type GetAppLogsOK struct {

	// In: body
	Payload io.ReadCloser `json:"body,omitempty"`
}

// NewGetAppLogsOK creates GetAppLogsOK with default headers values
func NewGetAppLogsOK() *GetAppLogsOK {
	return &GetAppLogsOK{}
}

// WithPayload adds the payload to the get app logs o k response
func (o *GetAppLogsOK) WithPayload(payload io.ReadCloser) *GetAppLogsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get app logs o k response
func (o *GetAppLogsOK) SetPayload(payload io.ReadCloser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAppLogsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetAppLogsDefault Unexpected error

swagger:response getAppLogsDefault
*/
type GetAppLogsDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAppLogsDefault creates GetAppLogsDefault with default headers values
func NewGetAppLogsDefault(code int) *GetAppLogsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetAppLogsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get app logs default response
func (o *GetAppLogsDefault) WithStatusCode(code int) *GetAppLogsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get app logs default response
func (o *GetAppLogsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get app logs default response
func (o *GetAppLogsDefault) WithPayload(payload *models.Error) *GetAppLogsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get app logs default response
func (o *GetAppLogsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAppLogsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
