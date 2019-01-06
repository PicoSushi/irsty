package irasutoya

import (
	_ "fmt"
	"net/http"
	_ "time"
	_ "strconv"
	_ "strings"

	_ "github.com/PuerkitoBio/goquery"
)

// Entry is a page from irasutoya.com.
// If Entry.IsSpecial is false, Entry has one or more Irasutoes.
type Entry struct {
	URL string `json:"url"`

	IsSpecial bool `json:"is_special`

	Irasutoes []Irasuto `json:"-"`
}

// UpdateIrastoes fetches and updates all irastoes included in the entry's page.
// This is shorthand of FetchIrastoes and setting them to Entry's attribute Irastoes.
func (entry *Entry) UpdateIrastoes() error {
	irasutoes, err := entry.FetchIrastoes()
	if err != nil {
		return err
	}
	entry.Irasutoes = irasutoes
	return nil
}

// FetchIrastoes fetches all irasutoes included in the entry's page.
//
// NOTE: This method does not change Entry itself.
// If you want to update, call Entry.UpdateIrastoes instead.
func (entry *Entry)FetchIrastoes() ([]Irasuto, error) {
	var err error

	url := entry.URL
	// trim trailing slash
	if string(url[len(url)-1]) == "/" {
		url = url[0 : len(url)-1]
	}

	res, err := http.Get(url)
	if err != nil {
		return entry.Irasutoes, err
	}
	defer res.Body.Close()
	// WIP
	return entry.Irasutoes, err
}
