package docs

import (
	"github.com/go-openapi/runtime"
	"net/http"
)

// GetDocOK Success
type GetDocOK struct {
}

// NewGetDocOK creates GetDocOK with default headers values
func NewGetDocOK() *GetDocOK {
	return &GetDocOK{}
}

// WriteResponse to the client
func (o *GetDocOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	html := `<!DOCTYPE html>
	<html>
	  <head>
		<title>Atomica Blog</title>
		<!-- needed for adaptive design -->
		<meta charset="utf-8"/>
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<link href="https://fonts.googleapis.com/css?family=Montserrat:300,400,700|Roboto:300,400,700" rel="stylesheet">

		<!--
		ReDoc doesn't change outer page styles
		-->
		<style>
		  body {
			margin: 0;
			padding: 0;
		  }
		</style>
	  </head>
	  <body>
		<redoc spec-url='/swagger.json'></redoc>
		<script src="https://cdn.jsdelivr.net/npm/redoc@next/bundles/redoc.standalone.js"> </script>
	  </body>
	</html>`

	rw.Header().Add("content-type", "text/html")
	if _, err := rw.Write([]byte(html)); err != nil {
		rw.Write([]byte(err.Error()))
		//panic(err) // let the recovery middleware deal with this
	}

}
