package route

import (
	"shortenUrl/reader/config/database"
	"time"
)

func findRouteByShortURL(shortURL string) (Route, string) {
	connection := database.OpenConnectionDatabase()
	defer connection.Close()
	
	routeEntity := connection.QueryRow("SELECT * FROM ROUTE WHERE NAM_SHORT = $1", shortURL)

	if routeEntity != nil {
		var id int
		var destination string
		var shortURL string
		var createdAt time.Time
	
		routeEntity.Scan(&id, &destination, &shortURL, &createdAt)
	
		return of(id, destination, shortURL, createdAt), ""
	}

	return Route{}, "Could not find with parameter " + shortURL
}

func save(destination string, shortURL string) {
	connection := database.OpenConnectionDatabase()
	defer connection.Close()

	query := "INSERT INTO ROUTE(NAM_DESTINARION, NAM_SHORT) VALUES ($1, $2)"

	createRouteQueryPrepared, err := connection.Prepare(query)
	defer createRouteQueryPrepared.Close()

	if err != nil {
		panic(err)
	}
	createRouteQueryPrepared.Exec(destination, shortURL)
}
