package irasutoya

import (
	"testing"
)

func TestSearch(t *testing.T) {
	sres, err := Search("イクラのお寿司")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	ikura := SearchResult{}
	ok := false
	for _, sr := range sres {
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
