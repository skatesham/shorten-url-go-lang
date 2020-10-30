package route

import "time"

// Route is the domain object for shorten url operation
type Route struct {
	ID int `json:"id"`
	Destination string `json:"destination"`
	Shorten     string `json:"shorten"`
	CreatedAt   time.Time `json:"createdAt"`
}

// Of factory method for build URL
func of(id int, destination string, shorten string, createdAt time.Time) Route {
	return Route{ID: id, Destination: destination, Shorten: shorten, CreatedAt: createdAt}
}
