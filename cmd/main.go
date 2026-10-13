package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("test")
	http.ListenAndServe(":4000", setupRoutes())

}
