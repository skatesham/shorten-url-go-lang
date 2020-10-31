package route

// GetRoute get business object route containing destionation url
func GetRoute(shortURL string) (Route, string) {
	return findRouteByShortURL(shortURL)
}
