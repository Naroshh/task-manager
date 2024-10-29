// handlers/home.go
package handlers

import (
	"html/template"
	"net/http"
)

// HomeHandler serves the main HTML page.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Failed to load page", http.StatusInternalServerError)
	}
}
