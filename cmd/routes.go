package main

import "net/http"

func setupRoutes(g *group) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/ws", g)

	return mux
}
