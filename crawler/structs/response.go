package structs

// ResponseData stores the data in from response
type ResponseData struct {
	Results []struct {
		ExhaustiveNbHits bool `json:"exhaustiveNbHits"`
		Hits             []struct {
			Geoloc struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"_geoloc"`
			// HighlightResult struct {
			// 	Geoloc struct {
			// 		Lat struct {
			// 			MatchLevel   string        `json:"matchLevel"`
			// 			MatchedWords []interface{} `json:"matchedWords"`
			// 			Value        string        `json:"value"`
			// 		} `json:"lat"`
			// 		Lng struct {
			// 			MatchLevel   string        `json:"matchLevel"`
			// 			MatchedWords []interface{} `json:"matchedWords"`
			// 			Value        string        `json:"value"`
			// 		} `json:"lng"`
			// 	} `json:"_geoloc"`
			// 	AreaName struct {
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 		Value        string        `json:"value"`
			// 	} `json:"area_name"`
			// 	CityName struct {
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 		Value        string        `json:"value"`
			// 	} `json:"city_name"`
			// 	CountryName struct {
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 		Value        string        `json:"value"`
			// 	} `json:"country_name"`
			// 	CuisineType struct {
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 		Value        string        `json:"value"`
			// 	} `json:"cuisine_type"`
			// 	MichelinAward struct {
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 		Value        string        `json:"value"`
			// 	} `json:"michelin_award"`
			// 	Name struct {
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 		Value        string        `json:"value"`
			// 	} `json:"name"`
			// 	Postcode struct {
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 		Value        string        `json:"value"`
			// 	} `json:"postcode"`
			// 	RegionName struct {
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 		Value        string        `json:"value"`
			// 	} `json:"region_name"`
			// 	Street struct {
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 		Value        string        `json:"value"`
			// 	} `json:"street"`
			// } `json:"_highlightResult"`
			CityName      string `json:"city_name"`
			CountryName   string `json:"country_name"`
			CuisineType   string `json:"cuisine_type"`
			MichelinAward string `json:"michelin_award"`
			Name          string `json:"name"`
			// ObjectID      string `json:"objectID"`
			Street        string `json:"street"`
		} `json:"hits"`
		// HitsPerPage       float64 `json:"hitsPerPage"`
		// Index             string  `json:"index"`
		// NbHits            float64 `json:"nbHits"`
		// NbPages           float64 `json:"nbPages"`
		// Page              float64 `json:"page"`
		// Params            string  `json:"params"`
		// ProcessingTimeMS  float64 `json:"processingTimeMS"`
		// Query             string  `json:"query"`
		// QueryAfterRemoval string  `json:"queryAfterRemoval"`
	}`json:"results"`
} 