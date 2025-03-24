package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Renders a page using base.tmpl.html and a specific page template
func (app *application) renderTemplate(w http.ResponseWriter, page string) {
    files := []string{
        "./ui/html/base.tmpl.html",      // Base template
        "./ui/html/pages/" + page + ".tmpl", // Dynamic page template
    }

    ts, err := template.ParseFiles(files...)
    if err != nil {
        app.serverError(w, err)
        return
    }

    err = ts.ExecuteTemplate(w, "base", nil)
    if err != nil {
        app.serverError(w, err)
    }
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.NotFound(w)
		return
	}

	app.renderTemplate(w, "home")
}

func (app *application) bugView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.NotFound(w)
		return
	}
	fmt.Fprintf(w, "Display bug with ID %d...", id)
}

func (app *application) bugCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new bug\n"))
}
