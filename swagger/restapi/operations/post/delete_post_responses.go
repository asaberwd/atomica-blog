// Code generated by go-swagger; DO NOT EDIT.

package post

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/asaberwd/atomica-blog/swagger/models"
)

// DeletePostOKCode is the HTTP code returned for type DeletePostOK
const DeletePostOKCode int = 200

/*DeletePostOK success

swagger:response deletePostOK
*/
type DeletePostOK struct {
}

// NewDeletePostOK creates DeletePostOK with default headers values
func NewDeletePostOK() *DeletePostOK {

	return &DeletePostOK{}
}

// WriteResponse to the client
func (o *DeletePostOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeletePostBadRequestCode is the HTTP code returned for type DeletePostBadRequest
const DeletePostBadRequestCode int = 400

/*DeletePostBadRequest Invalid request

swagger:response deletePostBadRequest
*/
type DeletePostBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeletePostBadRequest creates DeletePostBadRequest with default headers values
func NewDeletePostBadRequest() *DeletePostBadRequest {

	return &DeletePostBadRequest{}
}

// WithPayload adds the payload to the delete post bad request response
func (o *DeletePostBadRequest) WithPayload(payload *models.Error) *DeletePostBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete post bad request response
func (o *DeletePostBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeletePostBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeletePostNotFoundCode is the HTTP code returned for type DeletePostNotFound
const DeletePostNotFoundCode int = 404

/*DeletePostNotFound Not found.

swagger:response deletePostNotFound
*/
type DeletePostNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeletePostNotFound creates DeletePostNotFound with default headers values
func NewDeletePostNotFound() *DeletePostNotFound {

	return &DeletePostNotFound{}
}

// WithPayload adds the payload to the delete post not found response
func (o *DeletePostNotFound) WithPayload(payload *models.Error) *DeletePostNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete post not found response
func (o *DeletePostNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeletePostNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
