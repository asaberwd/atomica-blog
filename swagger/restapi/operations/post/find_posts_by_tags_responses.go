// Code generated by go-swagger; DO NOT EDIT.

package post

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/asaberwd/atomica-blog/swagger/models"
)

// FindPostsByTagsOKCode is the HTTP code returned for type FindPostsByTagsOK
const FindPostsByTagsOKCode int = 200

/*FindPostsByTagsOK successful operation

swagger:response findPostsByTagsOK
*/
type FindPostsByTagsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Post `json:"body,omitempty"`
}

// NewFindPostsByTagsOK creates FindPostsByTagsOK with default headers values
func NewFindPostsByTagsOK() *FindPostsByTagsOK {

	return &FindPostsByTagsOK{}
}

// WithPayload adds the payload to the find posts by tags o k response
func (o *FindPostsByTagsOK) WithPayload(payload []*models.Post) *FindPostsByTagsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find posts by tags o k response
func (o *FindPostsByTagsOK) SetPayload(payload []*models.Post) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindPostsByTagsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Post, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// FindPostsByTagsBadRequestCode is the HTTP code returned for type FindPostsByTagsBadRequest
const FindPostsByTagsBadRequestCode int = 400

/*FindPostsByTagsBadRequest Invalid tag value

swagger:response findPostsByTagsBadRequest
*/
type FindPostsByTagsBadRequest struct {
}

// NewFindPostsByTagsBadRequest creates FindPostsByTagsBadRequest with default headers values
func NewFindPostsByTagsBadRequest() *FindPostsByTagsBadRequest {

	return &FindPostsByTagsBadRequest{}
}

// WriteResponse to the client
func (o *FindPostsByTagsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}
