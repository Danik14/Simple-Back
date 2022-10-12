package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Danik14/simpleBack/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	//making path exact
	//if leave it without it
	//any url string will be sufficient to activate this handler
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	// for _, snippet := range snippets {
	// 	fmt.Fprintf(w, "%+v\n\n\n", snippet)
	// }

	// Call the newTemplateData() helper to get a templateData struct containing // the 'default' data (which for now is just the current year), and add the // snippets slice to it.
	data := app.newTemplateData(r)
	data.Snippets = snippets

	app.render(w, http.StatusOK, "home.tmpl.html", data)
}

// Add a snippetView handler function.
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {

	//validating id string, because it is entered by user
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	// Use the SnippetModel object's Get method to retrieve the data for a
	// specific record based on its ID. If no matching record is found,
	// return a 404 Not Found response.
	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}

		return
	}

	// Call the newTemplateData() helper to get a templateData struct containing // the 'default' data (which for now is just the current year), and add the // snippets slice to it.
	data := app.newTemplateData(r)
	data.Snippet = snippet

	app.render(w, http.StatusOK, "view.tmpl.html", data)
}

// Add a snippetCreate handler function.
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {

	//making this handler to only allow post requests
	if r.Method != "POST" {
		w.Header().Set("Allowsf", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "Some data"
	content := "Some more data\nTotototo UWU UWU\n LET'S GO"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}
