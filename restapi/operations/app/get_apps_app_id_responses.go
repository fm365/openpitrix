// Code generated by go-swagger; DO NOT EDIT.

package app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"AppHub/models"
)

// GetAppsAppIDOKCode is the HTTP code returned for type GetAppsAppIDOK
const GetAppsAppIDOKCode int = 200

/*GetAppsAppIDOK An app

swagger:response getAppsAppIdOK
*/
type GetAppsAppIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.App `json:"body,omitempty"`
}

// NewGetAppsAppIDOK creates GetAppsAppIDOK with default headers values
func NewGetAppsAppIDOK() *GetAppsAppIDOK {
	return &GetAppsAppIDOK{}
}

// WithPayload adds the payload to the get apps app Id o k response
func (o *GetAppsAppIDOK) WithPayload(payload *models.App) *GetAppsAppIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get apps app Id o k response
func (o *GetAppsAppIDOK) SetPayload(payload *models.App) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAppsAppIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAppsAppIDNotFoundCode is the HTTP code returned for type GetAppsAppIDNotFound
const GetAppsAppIDNotFoundCode int = 404

/*GetAppsAppIDNotFound The app does not exist.

swagger:response getAppsAppIdNotFound
*/
type GetAppsAppIDNotFound struct {
}

// NewGetAppsAppIDNotFound creates GetAppsAppIDNotFound with default headers values
func NewGetAppsAppIDNotFound() *GetAppsAppIDNotFound {
	return &GetAppsAppIDNotFound{}
}

// WriteResponse to the client
func (o *GetAppsAppIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

// GetAppsAppIDInternalServerErrorCode is the HTTP code returned for type GetAppsAppIDInternalServerError
const GetAppsAppIDInternalServerErrorCode int = 500

/*GetAppsAppIDInternalServerError An unexpected error occured.

swagger:response getAppsAppIdInternalServerError
*/
type GetAppsAppIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAppsAppIDInternalServerError creates GetAppsAppIDInternalServerError with default headers values
func NewGetAppsAppIDInternalServerError() *GetAppsAppIDInternalServerError {
	return &GetAppsAppIDInternalServerError{}
}

// WithPayload adds the payload to the get apps app Id internal server error response
func (o *GetAppsAppIDInternalServerError) WithPayload(payload *models.Error) *GetAppsAppIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get apps app Id internal server error response
func (o *GetAppsAppIDInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAppsAppIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}