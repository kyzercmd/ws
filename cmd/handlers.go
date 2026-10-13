package main

import "net/http"

func serveFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./video.mp4")
}
