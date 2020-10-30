package interfaces

import "net/http"

// LoadRoutes map all route of application
func LoadRoutes() {
	http.HandleFunc("/", MappingRouteShortURL)
}

