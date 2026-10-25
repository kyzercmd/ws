package main

import "net/http"

func (g *group) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := &client{
		socket:  conn,
		receive: make(chan []byte, 256),
		group:   g,
	}

	g.join <- client

	defer func() {
		g.leave <- client
	}()

	go client.write()
	client.read()
}
