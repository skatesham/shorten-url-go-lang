package interfaces

import (
	"net/http"
)

// MappingRouteShortURL mapping route for shorten url
func MappingRouteShortURL(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		GetMappingReadRouteShortURL(w, r)
	} else if r.Method == "POST" {
		PostMappingCreateRouteShortURL(w, r)
	}
}
