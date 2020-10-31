package route

import (
	"net/url"

	"github.com/google/uuid"
)

// CreateRoute Create a new route
func CreateRoute(destination string) (Route, string) {
	if !isURL(destination) {
		return Route{}, "Must be an URL."
	}

	shortURL := generateUUID()
	save(destination, shortURL)

	return findRouteByShortURL(shortURL)
}

func isURL(str string) bool {
    u, err := url.Parse(str)
    return err == nil && u.Scheme != "" && u.Host != ""
}

func generateUUID() string {
	uuid, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	return uuid.String()
}