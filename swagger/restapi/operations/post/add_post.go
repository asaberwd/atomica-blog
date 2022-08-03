// Code generated by go-swagger; DO NOT EDIT.

package post

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AddPostHandlerFunc turns a function with the right signature into a add post handler
type AddPostHandlerFunc func(AddPostParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddPostHandlerFunc) Handle(params AddPostParams) middleware.Responder {
	return fn(params)
}

// AddPostHandler interface for that can handle valid add post params
type AddPostHandler interface {
	Handle(AddPostParams) middleware.Responder
}

// NewAddPost creates a new http.Handler for the add post operation
func NewAddPost(ctx *middleware.Context, handler AddPostHandler) *AddPost {
	return &AddPost{Context: ctx, Handler: handler}
}

/*AddPost swagger:route POST /posts post addPost

Add a new post to the blog

*/
type AddPost struct {
	Context *middleware.Context
	Handler AddPostHandler
}

func (o *AddPost) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddPostParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
