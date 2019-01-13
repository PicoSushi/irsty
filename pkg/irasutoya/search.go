package irasutoya

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// SearchResult is result of search.
// Search returns multiple SearchResult.
// Result will corresponds Entry one-to-one.
type SearchResult struct {
	EntryURL     string `json:"entry_url"`
	ThumbnailURL string `json:"thumbnail_url"`
	Description  string `json:"description"`
}

// Search searches irasuto from irasutoya with given query, and returns []SearchResult.
func Search(query string) ([]SearchResult, error) {
	searchURL := fmt.Sprintf("https://www.irasutoya.com/search?q=%s", query)
	searchResults, nextURL, err := fetchSearchResult(searchURL)

	if err != nil {
		return searchResults, err
	}

	for nextURL != "" {
		srs, n, err := fetchSearchResult(nextURL)
		searchResults = append(searchResults, srs...)
		if err != nil {
			return searchResults, err
		}
		nextURL = n
	}

	return searchResults, err
}

// fetchSearchResult fetches and parses URL.
func fetchSearchResult(url string) (srs []SearchResult, nextURL string, err error) {
	res, err := http.Get(url)
	if err != nil {
		return srs, "", err
	}
	defer res.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return srs, "", err
	}

	posts := doc.Find("#post") // wow, irasutoya has multiple ids!

	posts.Each(func(i int, s *goquery.Selection) {
		sr := SearchResult{}
		desc := s.Find("div.boxmeta.clearfix > h2 > a")
		sr.EntryURL, _ = desc.Attr("href")
		sr.Description = desc.Text()

		script := s.Find("div.boxim > a > script").Text()
		t := strings.Split(script, "\"")
		sr.ThumbnailURL = t[1]

		srs = append(srs, sr)
	})
	nextURL, _ = doc.Find("#Blog1_blog-pager-older-link").Attr("href")

	return srs, nextURL, nil
}
