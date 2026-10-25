package main

import "net/http"

func setupRoutes(g *group) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/ws", g)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("/", homeHandler)

	return mux
}
