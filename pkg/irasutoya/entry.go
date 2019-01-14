package irasutoya

import (
	"fmt"
	"net/http"
	_ "strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Entry is a page from irasutoya.com.
// If Entry.IsSpecial is false, Entry has one or more Irasutoes.
type Entry struct {
	URL string `json:"url"`

	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsSpecial   bool      `json:"is_special"` // ex: PR page
	PublishDate time.Time `json:"publish_date"`

	Categories []string `json:"categories"`
	Irasutoes []Irasuto `json:"irasutoes"`
}

// NewEntry is a constructor for Entry.
func NewEntry(url string) (Entry, error) {
	entry := Entry{
		URL: url,
	}
	err := entry.Load()
	return entry, err
}

// Load fetches and updates all irastoes included in the entry's page.
// This is shorthand of FetchIrastoes and setting them to Entry's attribute Irastoes.
func (entry *Entry) Load() error {
	var err error
	url := entry.URL
	if entry.URL == "" {
		return fmt.Errorf("entry.url not set.")
	}

	// trim trailing slash
	if string(url[len(url)-1]) == "/" {
		url = url[0 : len(url)-1]
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	entry.Title = strings.Trim(doc.Find("#post > div.title > h2").Text(), " \n")
	entry.Description = strings.Trim(doc.Find("#post > div.entry > div").Text(), " \n")

	pubDate := doc.Find("#post > div:nth-child(3) > div.entry-post-date > span").Text()
	loc, _ := time.LoadLocation("Asia/Tokyo")
	entry.PublishDate, err = time.ParseInLocation("公開日：2006/01/02", pubDate, loc)
	if err != nil {
		return err
	}
	return nil
}
