package main

import "net/http"

func setupRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/watch", serveFile)

	return mux
}
