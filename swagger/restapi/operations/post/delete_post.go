// Code generated by go-swagger; DO NOT EDIT.

package post

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeletePostHandlerFunc turns a function with the right signature into a delete post handler
type DeletePostHandlerFunc func(DeletePostParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeletePostHandlerFunc) Handle(params DeletePostParams) middleware.Responder {
	return fn(params)
}

// DeletePostHandler interface for that can handle valid delete post params
type DeletePostHandler interface {
	Handle(DeletePostParams) middleware.Responder
}

// NewDeletePost creates a new http.Handler for the delete post operation
func NewDeletePost(ctx *middleware.Context, handler DeletePostHandler) *DeletePost {
	return &DeletePost{Context: ctx, Handler: handler}
}

/*DeletePost swagger:route DELETE /post/{postId} post deletePost

Soft Delete a post

*/
type DeletePost struct {
	Context *middleware.Context
	Handler DeletePostHandler
}

func (o *DeletePost) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeletePostParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
