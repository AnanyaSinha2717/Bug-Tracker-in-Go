package main

import(
	"fmt"
	"net/http"
	"runtime/debug"
)

// The serverError helper writes an error message and stack trace (execution path) to the errorLog,
// then sends a generic 500 Internal Server Error response to the user.
func (app *application) serverError(w http.ResponseWriter, err error){
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errlog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// sends to user
func (app *application) clientError(w http.ResponseWriter, status int){
	http.Error(w, http.StatusText(status), status)
}

func (app *application) NotFound(w http.ResponseWriter){
	app.clientError(w, http.StatusNotFound)
}