package route

import (
	"shortenUrl/reader/config/database"
	"time"
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
