package main

import (
	"go-htmx-todo/handlers"
	"net/http"
)

func main() {

	http.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.Handle("GET /", handlers.HTTPHandler(handlers.HomeHandler))

	http.ListenAndServe(":3000", nil)
}
