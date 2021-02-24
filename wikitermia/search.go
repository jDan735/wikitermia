package wikitermia

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"log"
)

type SearchResults struct {
	Title     string
	Pageid    int
	Wordcount int
	Snippet   string
	Timestamp string
}

type SearchResponse struct {
	Query struct {
		Search []SearchResults
	}
}

func Search(query string, limit string) (results []SearchResults){
	url := HOST + "action=query&format=json&list=search&srlimit=" + limit + "&srsearch=" + query

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	response := SearchResponse{}
	json.Unmarshal(body, &response)

	results = response.Query.Search

	return
}
