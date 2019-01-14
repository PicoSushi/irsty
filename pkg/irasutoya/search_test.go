package irasutoya

import (
	"testing"
)

func TestSearch(t *testing.T) {
	srs, err := Search("イクラのお寿司")
	if err != nil {
		t.Fatalf("Error Search()ing: %#v", err)
	}

	ikura := SearchResult{}
	ok := false
	for _, sr := range srs {
		if sr.Description == "イクラのお寿司のイラスト" {
			ok = true
			ikura = sr
		}
	}
	if !ok {
		t.Fatalf("Couldn't find イクラのお寿司のイラスト.")
	}
	if ikura.EntryURL != "https://www.irasutoya.com/2013/04/blog-post_8611.html" {
		t.Fatalf("EntryURL is odd. changed?")
	}
	if ikura.ThumbnailURL == "" {
		t.Fatalf("ThumbnailURL not set.")
	}
}

func TestTwoResultsSearch(t *testing.T) {
	srs, err := Search("卵焼き")
	if err != nil {
		t.Fatalf("Error Search()ing: %#v", err)
	}

	datemaki := SearchResult{}
	ok := false
	for _, sr := range srs {
		if sr.Description == "伊達巻のイラスト（おせち料理）" {
			ok = true
			datemaki = sr
		}
	}
	if !ok {
		t.Fatalf("Couldn't find 伊達巻のイラスト（おせち料理）.")
	}
	if datemaki.EntryURL != "https://www.irasutoya.com/2015/11/blog-post_653.html" {
		t.Fatalf("EntryURL is odd. changed?")
	}
	if datemaki.ThumbnailURL == "" {
		t.Fatalf("ThumbnailURL not set.")
	}
}

func TestManyResultsSerach(t *testing.T) {
	srs, err := Search("忍者")
	if err != nil {
		t.Fatalf("Error Search()ing: %#v", err)
	}
	t.Logf("len(Search(\"忍者\")) has %d.", len(srs))
	if len(srs) < 30 {
		t.Fatalf("Searching \"忍者\"'s result too less: %d.", len(srs))
	}
}

func TestFetchEntry(t *testing.T) {
	// イクラのお寿司
	ikuraURL := "https://www.irasutoya.com/2013/04/blog-post_8611.html"
	sr := SearchResult{EntryURL: ikuraURL}
	entry, err := sr.FetchEntry()
	if err != nil {
		t.Fatal(err)
	}

	if entry.URL != ikuraURL {
		t.Fatalf("URL changed: \"%s\", expected: \"%s\"", entry.URL, ikuraURL)
	}

	if entry.Title != "イクラのお寿司のイラスト" {
		t.Fatalf("Title seems odd: %s", entry.Title)
	}

	if len(entry.Irasutoes) != 1 {
		t.Fatalf("Number of Ikura is odd: %d, want: 1", len(entry.Irasutoes))
	}

	ikuraIrasuto := entry.Irasutoes[0]
	if ikuraIrasuto.Title != "イクラのお寿司のイラスト" {
		t.Fatalf("Irasuto title seems odd: %s", ikuraIrasuto.Title)
	}

	if ikuraIrasuto.ImageURL == "" {
		t.Fatalf("Irasuto.ImageURL not set.")
	}
}
