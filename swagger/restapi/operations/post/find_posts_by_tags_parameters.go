// Code generated by go-swagger; DO NOT EDIT.

package post

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewFindPostsByTagsParams creates a new FindPostsByTagsParams object
//
// There are no default values defined in the spec.
func NewFindPostsByTagsParams() FindPostsByTagsParams {

	return FindPostsByTagsParams{}
}

// FindPostsByTagsParams contains all the bound params for the find posts by tags operation
// typically these are obtained from a http.Request
//
// swagger:parameters findPostsByTags
type FindPostsByTagsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Tags to filter by
	  Required: true
	  In: query
	  Collection Format: multi
	*/
	Tags []string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewFindPostsByTagsParams() beforehand.
func (o *FindPostsByTagsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qTags, qhkTags, _ := qs.GetOK("tags")
	if err := o.bindTags(qTags, qhkTags, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindTags binds and validates array parameter Tags from query.
//
// Arrays are parsed according to CollectionFormat: "multi" (defaults to "csv" when empty).
func (o *FindPostsByTagsParams) bindTags(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("tags", "query", rawData)
	}
	// CollectionFormat: multi
	tagsIC := rawData
	if len(tagsIC) == 0 {
		return errors.Required("tags", "query", tagsIC)
	}

	var tagsIR []string
	for _, tagsIV := range tagsIC {
		tagsI := tagsIV

		tagsIR = append(tagsIR, tagsI)
	}

	o.Tags = tagsIR

	return nil
}
