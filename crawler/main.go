package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mgcrawler/structs"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/levigross/grequests"
)

const (
	maxNumPerPage       int    = 1000
	totalNumRestaurants int    = 13821
	sleepDuration       int    = 1000
	indexName           string = "prod-restaurants-en"
)

var (
	params = map[string]string{
		"x-algolia-agent":          "Algolia for JavaScript (3.35.1); Browser (lite); instantsearch.js (4.1.1); JS Helper (3.0.0)",
		"x-algolia-application-id": "8NVHRD7ONV",
		"x-algolia-api-key":        "71b3cff102a474b924dfcb9897cc6fa8",
	}
	queryAttributes = []string{
		"_geoloc",
		"city_name",
		"country_name",
		"cuisine_type",
		"street",
		"michelin_award",
		"name",
	}
)

func main() {
	_, es := os.Stat("../data/")
	if es != nil && os.IsNotExist(es) {
		os.MkdirAll("../data", os.ModePerm)
	}
	wg := &sync.WaitGroup{}
	for i := 0; i < totalNumRestaurants/maxNumPerPage+1; i++ {
		wg.Add(1)
		go crawl(i, wg)
		time.Sleep(time.Second * 1)
	}
	wg.Wait()
}

func crawl(pageIndex int, wg *sync.WaitGroup) {

	fmt.Println("[LOG] Start crawling page", pageIndex)
	rp := structs.RequestParams{
		Page:           pageIndex,
		Attributes:     queryAttributes,
		NumbersPerPage: maxNumPerPage,
	}
	ro := &grequests.RequestOptions{
		Params: params,
		JSON:   rp.JSON(indexName),
	}
	resp, err := grequests.Post("https://8nvhrd7onv-2.algolianet.com/1/indexes/*/queries", ro)
	if err != nil {
		fmt.Println(err)
		return
	}

	res := structs.ResponseData{}
	resp.JSON(&res)
	bytes, errJSON := json.Marshal(res.Results[0])
	if errJSON != nil {
		fmt.Printf("[ERROR] %s", errJSON.Error())
	} else {
		filename := filepath.Join("../data", fmt.Sprintf("data%d.json", pageIndex))
		ioutil.WriteFile(filename, bytes, 0644)
	}
	fmt.Println("[LOG]Finished crawling page", pageIndex)
	wg.Done()
}
