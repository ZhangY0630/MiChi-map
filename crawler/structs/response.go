package structs

// RestaurantInfo holds the fields that is required for crawled data
type RestaurantInfo struct {
	CityName      string `json:"city_name"`
	CountryName   string `json:"country_name"`
	CuisineType   string `json:"cuisine_type"`
	MichelinAward string `json:"michelin_award"`
	Street        string `json:"street"`
}

// RespResult is used for holding the original request
type RespResult struct {
	Restaurants []RestaurantInfo `json:"hits"`
}
