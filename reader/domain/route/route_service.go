package route

import (
	"net/url"
	"shortenUrl/reader/config/database"
	"time"

	"github.com/google/uuid"
)

// GetRoute get business object route containing destionation url
func GetRoute(shortURL string) (Route, string) {
	connection := database.OpenConnectionDatabase()
	defer connection.Close()
	
	routeEntity, err := connection.Query("SELECT * FROM ROUTE WHERE NAM_SHORT = $1", shortURL)
	defer routeEntity.Close()

	if err != nil {
		panic(err)
	}

	if routeEntity.Next() {
		var id int
		var destination string
		var shortURL string
		var createdAt time.Time
	
		routeEntity.Scan(&id, &destination, &shortURL, &createdAt)
	
		return of(id, destination, shortURL, createdAt), ""
	}

	return Route{}, "Could not find with parameter " + shortURL
}


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