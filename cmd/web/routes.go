package main

import "net/http"

func (app *application) routes(static_dir *string) *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(*static_dir))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	return mux
}