package handlers

import (
	"net/http"

	"go-htmx-todo/views"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) error {
	return render(w, r, views.Home())
}
