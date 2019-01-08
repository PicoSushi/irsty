package irasutoya

import (
	"testing"
)

func TestSearch(t *testing.T) {
	sres, err := Search("イクラのお寿司")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	ok := false
	for _, sr := range sres {
		if sr.Description == "イクラのお寿司のイラスト" {
			ok = true
		}
	}
	if !ok {
		t.Fatalf("Couldn't find イクラのお寿司のイラスト.")
	}
}
