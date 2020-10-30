package interfaces

import (
	"encoding/json"
	"net/http"
	"shortenUrl/reader/domain/route"
)

// PostMappingCreateRouteShortURL create an short url route
func PostMappingCreateRouteShortURL(w http.ResponseWriter, r *http.Request) {

	var routeVO route.Route

	err := json.NewDecoder(r.Body).Decode(&routeVO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	found, errors := route.CreateRoute(routeVO.Destination)

	if errors != "" {
		w.Header().Add("error", errors)
		w.WriteHeader(http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(found)
}

