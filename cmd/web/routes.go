package main

import (
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes(static_dir *string) http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(*static_dir))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	dynamic := alice.New(app.sessionManager.LoadAndSave)

	mux.Handle("GET /{$}", dynamic.ThenFunc(app.home))
	mux.Handle("GET /snippet/view/{id}", dynamic.ThenFunc(app.snippetView))
	mux.Handle("GET /snippet/create", dynamic.ThenFunc(app.snippetCreate))
	mux.Handle("POST /snippet/create", dynamic.ThenFunc(app.snippetCreatePost))

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}
