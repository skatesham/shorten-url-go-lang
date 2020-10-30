package main

import (
	"net/http"
	"shortenUrl/reader/interfaces"
)

func main() {
	interfaces.LoadRoutes()
	http.ListenAndServe(":8080", nil)
}