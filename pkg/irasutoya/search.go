package irasutoya

import (
	_ "fmt"
	_ "net/http"
)

type SearchResult struct {
	URL string `json:"url"`
	ThumbnailURL string `json:"thumbnail_url"`
	Description string `json:"description"`
}

// Search searches irasuto from irasutoya with given query, and returns []SearchResult.
func Search(quuery string) ([]SearchResult, error) {
	var searchResults = []SearchResult{}
	return searchResults, nil
}

// fetchSearchResult fetches and parses URL.
func fetchSearchResult(url string) (sres []SearchResult, hasNext bool, err error) {
	return sres, false, nil
}
