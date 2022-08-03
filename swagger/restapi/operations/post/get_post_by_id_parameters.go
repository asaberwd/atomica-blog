// Code generated by go-swagger; DO NOT EDIT.

package post

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPostByIDParams creates a new GetPostByIDParams object
// no default values defined in spec.
func NewGetPostByIDParams() GetPostByIDParams {

	return GetPostByIDParams{}
}

// GetPostByIDParams contains all the bound params for the get post by Id operation
// typically these are obtained from a http.Request
//
// swagger:parameters getPostById
type GetPostByIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*ID of post to return
	  Required: true
	  In: path
	*/
	PostID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetPostByIDParams() beforehand.
func (o *GetPostByIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rPostID, rhkPostID, _ := route.Params.GetOK("postId")
	if err := o.bindPostID(rPostID, rhkPostID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindPostID binds and validates parameter PostID from path.
func (o *GetPostByIDParams) bindPostID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("postId", "path", "int64", raw)
	}
	o.PostID = value

	return nil
}