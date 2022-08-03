// Code generated by go-swagger; DO NOT EDIT.

package comments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AddPostCommentHandlerFunc turns a function with the right signature into a add post comment handler
type AddPostCommentHandlerFunc func(AddPostCommentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddPostCommentHandlerFunc) Handle(params AddPostCommentParams) middleware.Responder {
	return fn(params)
}

// AddPostCommentHandler interface for that can handle valid add post comment params
type AddPostCommentHandler interface {
	Handle(AddPostCommentParams) middleware.Responder
}

// NewAddPostComment creates a new http.Handler for the add post comment operation
func NewAddPostComment(ctx *middleware.Context, handler AddPostCommentHandler) *AddPostComment {
	return &AddPostComment{Context: ctx, Handler: handler}
}

/*AddPostComment swagger:route POST /post/{postId}/comments comments addPostComment

Add a new comment to a post

*/
type AddPostComment struct {
	Context *middleware.Context
	Handler AddPostCommentHandler
}

func (o *AddPostComment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddPostCommentParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
