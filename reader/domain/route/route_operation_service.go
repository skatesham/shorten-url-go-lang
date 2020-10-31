package route

import (
	"net/url"
	"shortenUrl/reader/config/database"

	"github.com/google/uuid"
)

// CreateRoute Create a new route
func CreateRoute(destination string) (Route, string) {
	if !isURL(destination) {
		return Route{}, "Must be an URL."
	}

	connection := database.OpenConnectionDatabase()
	defer connection.Close()

	query := "INSERT INTO ROUTE(NAM_DESTINARION, NAM_SHORT) VALUES ($1, $2)"

	createRouteQueryPrepared, err := connection.Prepare(query)

	defer createRouteQueryPrepared.Close()

	if err != nil {
		panic(err)
	}

	shortURL := generateUUID()
	createRouteQueryPrepared.Exec(destination, shortURL)

	return GetRoute(shortURL)
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