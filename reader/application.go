package main

import (
	"net/http"
	"shortenUrl/reader/config"
	"shortenUrl/reader/interfaces"
)

func main() {
	config.SetUpProfileConfiguration()
	interfaces.LoadRoutes()
	http.ListenAndServe(":8080", nil)
}