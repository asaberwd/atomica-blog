// Code generated by go-swagger; DO NOT EDIT.

package post

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/asaberwd/atomica-blog/swagger/models"
)

// AddPostOKCode is the HTTP code returned for type AddPostOK
const AddPostOKCode int = 200

/*AddPostOK successful operation

swagger:response addPostOK
*/
type AddPostOK struct {

	/*
	  In: Body
	*/
	Payload *models.Post `json:"body,omitempty"`
}

// NewAddPostOK creates AddPostOK with default headers values
func NewAddPostOK() *AddPostOK {

	return &AddPostOK{}
}

// WithPayload adds the payload to the add post o k response
func (o *AddPostOK) WithPayload(payload *models.Post) *AddPostOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add post o k response
func (o *AddPostOK) SetPayload(payload *models.Post) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddPostOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddPostBadRequestCode is the HTTP code returned for type AddPostBadRequest
const AddPostBadRequestCode int = 400

/*AddPostBadRequest Invalid request

swagger:response addPostBadRequest
*/
type AddPostBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddPostBadRequest creates AddPostBadRequest with default headers values
func NewAddPostBadRequest() *AddPostBadRequest {

	return &AddPostBadRequest{}
}

// WithPayload adds the payload to the add post bad request response
func (o *AddPostBadRequest) WithPayload(payload *models.Error) *AddPostBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add post bad request response
func (o *AddPostBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddPostBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
