package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	//making path exact
	//if leave it without it
	//any url string will be sufficient to activate this handler
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	// Initialize a slice containing the paths to the two files. It's important // to note that the file containing our base template must be the *first* // file in the slice.
	files := []string{
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/base.tmpl.html",
		"./ui/html/pages/home.tmpl.html"}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
	// Use the ExecuteTemplate() method to write the content of the "base" // template as the response body.
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
		http.Error(w, "Internal Server Error", 500)
	}
}

// Add a snippetView handler function.
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {

	//validating id string, because it is entered by user
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Snipped with id value of %v", id)
}

// Add a snippetCreate handler function.
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {

	//making this handler to only allow post requests
	if r.Method != "POST" {
		w.Header().Set("Allowsf", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}
