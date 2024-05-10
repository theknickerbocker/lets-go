package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":7777", "HTTP network address")
	static_dir := flag.String("static", "ui/static/", "Path to static assets")
	flag.Parse()

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(*static_dir))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("starting server on ", *addr)

	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
