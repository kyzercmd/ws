package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	g := newGroup()

	fmt.Println("Starting")

	go g.run()
	if err := http.ListenAndServe(":4000", setupRoutes(g)); err != nil {
		log.Fatal(err)
	}

}
