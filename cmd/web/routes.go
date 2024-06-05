package main

import (
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes(static_dir *string) http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(*static_dir))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}
