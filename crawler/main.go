package main

import (
	"fmt"
	"mgcrawler/structs"

	"github.com/levigross/grequests"
)

func main() {
	rp := structs.RequestParams{
		Page: 0,
		Attributes: []string{
			"_geoloc",
			"city_name",
			"country_name",
			"cuisine_type",
			"street",
			"michelin_award",
			"name",
			"url",
		},
		NumbersPerPage: 1,
	}
	fmt.Println(string(rp.JSON("prod-restaurants-en")))

	ro := &grequests.RequestOptions{
		Params: map[string]string{
			"x-algolia-agent":          "Algolia for JavaScript (3.35.1); Browser (lite); instantsearch.js (4.1.1); JS Helper (3.0.0)",
			"x-algolia-application-id": "8NVHRD7ONV",
			"x-algolia-api-key":        "71b3cff102a474b924dfcb9897cc6fa8",
		},
		JSON: rp.JSON("prod-restaurants-en"),
	}
	resp, err := grequests.Post("https://8nvhrd7onv-2.algolianet.com/1/indexes/*/queries", ro)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp.String())
		resp.DownloadToFile("page1.json")
		res := structs.RespResult{}
		resp.JSON(&res)
		fmt.Println(len(res.Restaurants))
	}
}
