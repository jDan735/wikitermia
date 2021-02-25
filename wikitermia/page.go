package wikitermia

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
    "strconv"
	"log"
//	"fmt"
)

type WikipediaPage struct {
	Extract string
	Title string
	Pageid int
	Ns int
}

type PageResponse struct {
	Query struct {
		Pages []WikipediaPage
	}
}

func Page(pageid int) (page WikipediaPage){
	// params := "action=parse&prop=text&section=0&format=json&formatversion=2&pageid=" + strconv.Itoa(pageid)
	params := "action=query&prop=extracts&format=json&exintro&explaintext&formatversion=2&pageids=" + strconv.Itoa(pageid)

	resp, err := http.Get(HOST + params)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	response := PageResponse{}
	json.Unmarshal(body, &response)

	page = response.Query.Pages[0]
	
	return
}
