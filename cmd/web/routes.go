package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/bug/view/", app.bugView)
	mux.HandleFunc("/bug/create/", app.bugCreate)

	return mux;
}
