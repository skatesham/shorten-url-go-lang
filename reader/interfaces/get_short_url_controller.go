package interfaces

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shortenUrl/reader/domain/route"
)

// GetMappingReadRouteShortURL get short url route by path variable
func GetMappingReadRouteShortURL(w http.ResponseWriter, r *http.Request) {
	pathVariable := r.URL.Path[1:]
	found, err := route.GetRoute(pathVariable)
	if err != "" {
		fmt.Println(err)
		w.WriteHeader(http.StatusNoContent)
	}
	json.NewEncoder(w).Encode(found)
	// Redirect to destination option
	// http.Redirect(w, r, found.Destination, http.StatusSeeOther)
}
