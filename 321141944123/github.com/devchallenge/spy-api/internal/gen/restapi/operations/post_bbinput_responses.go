// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/devchallenge/spy-api/internal/gen/models"
)

// PostBbinputOKCode is the HTTP code returned for type PostBbinputOK
const PostBbinputOKCode int = 200

/*PostBbinputOK OK

swagger:response postBbinputOK
*/
type PostBbinputOK struct {
}

// NewPostBbinputOK creates PostBbinputOK with default headers values
func NewPostBbinputOK() *PostBbinputOK {

	return &PostBbinputOK{}
}

// WriteResponse to the client
func (o *PostBbinputOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PostBbinputBadRequestCode is the HTTP code returned for type PostBbinputBadRequest
const PostBbinputBadRequestCode int = 400

/*PostBbinputBadRequest Invalid arguments

swagger:response postBbinputBadRequest
*/
type PostBbinputBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostBbinputBadRequest creates PostBbinputBadRequest with default headers values
func NewPostBbinputBadRequest() *PostBbinputBadRequest {

	return &PostBbinputBadRequest{}
}

// WithPayload adds the payload to the post bbinput bad request response
func (o *PostBbinputBadRequest) WithPayload(payload *models.Error) *PostBbinputBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post bbinput bad request response
func (o *PostBbinputBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostBbinputBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostBbinputInternalServerErrorCode is the HTTP code returned for type PostBbinputInternalServerError
const PostBbinputInternalServerErrorCode int = 500

/*PostBbinputInternalServerError General server error

swagger:response postBbinputInternalServerError
*/
type PostBbinputInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostBbinputInternalServerError creates PostBbinputInternalServerError with default headers values
func NewPostBbinputInternalServerError() *PostBbinputInternalServerError {

	return &PostBbinputInternalServerError{}
}

// WithPayload adds the payload to the post bbinput internal server error response
func (o *PostBbinputInternalServerError) WithPayload(payload *models.Error) *PostBbinputInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post bbinput internal server error response
func (o *PostBbinputInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostBbinputInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
