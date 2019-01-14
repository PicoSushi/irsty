package irasutoya

import (
	"testing"
	"time"
)

func TestNewEntry(t *testing.T) {
	url := "https://www.irasutoya.com/2013/04/blog-post_8611.html"

	entry, err := NewEntry(url)
	if err != nil {
		t.Fatal("Error occurred on NewEntry:", err)
	}
	if entry.URL != url {
		t.Fatal(
			"URL has changed:", entry.URL,
			", expected:", url,
		)
	}
	exTitle := "イクラのお寿司のイラスト"
	if entry.Title != exTitle {
		t.Fatal(
			"Entry.Title seems odd:", entry.Title,
			", expected:", exTitle,
		)
	}
	exDesc := "ピカピカに光る美味しそうなイクラの軍艦巻きのイラストです。"
	if entry.Description != exDesc {
		t.Fatal(
			"Entry.Description seems odd:", entry.Description,
			", expected:", exDesc,
		)
	}

	if entry.IsSpecial {
		t.Fatal("This entry is set as Special, in spite of it's not special.")
	}
	year, month, day :=  entry.PublishDate.Date()
	if year != 2013 || month != time.April || day != 9 {
		t.Fatalf(
			"Entry.PublishDate seems odd: %d/%d/%d, expected: 2013/4/9",
			year, month, day,
		)
	}
}

func TestNewEntryWithMultipleIrastoes(t *testing.T) {
	url := "https://www.irasutoya.com/2018/03/blog-post_58.html"

	entry, err := NewEntry(url)
	if err != nil {
		t.Fatal("Error occurred on NewEntry:", err)
	}
	if entry.URL != url {
		t.Fatal(
			"URL has changed:", entry.URL,
			", expected:", url,
		)
	}
	exTitle := "いろいろな髪の色の女の子のイラスト"
	if entry.Title != exTitle {
		t.Fatal(
			"Entry.Title seems odd:", entry.Title,
			", expected:", exTitle,
		)
	}
	exDesc := "アニメやゲームの中に出てくるような、白、赤、オレンジ、黄色、緑、青、紫、茶色、ピンクなど様々な色の髪の毛の色をした女の子のキャラクターです。"
	if entry.Description != exDesc {
		t.Fatal(
			"Entry.Description seems odd:", entry.Description,
			", expected:", exDesc,
		)
	}

	if entry.IsSpecial {
		t.Fatal("This entry is set as Special, in spite of it's not special.")
	}
	year, month, day :=  entry.PublishDate.Date()
	if year != 2018 || month != time.March || day != 5 {
		t.Fatalf(
			"Entry.PublishDate seems odd: %d/%d/%d, expected: 2018/3/5",
			year, month, day,
		)
	}
}

func TestNewSpecialEntry(t *testing.T) {
	url := "https://www.irasutoya.com/2016/06/line.html"
	entry, err := NewEntry(url)
	if err != nil {
		t.Fatal("Error occurred on NewEntry:", err)
	}
	exTitle := "「いらすとやパーティ」がLINEスタンプになりました"
	if entry.Title != exTitle {
		t.Fatal(
			"Entry.Title seems odd:", entry.Title,
			", expected:", exTitle,
		)
	}
	if !entry.IsSpecial {
		t.Fatal("This entry is not set as Special, in spite of it's special.")
	}
	year, month, day :=  entry.PublishDate.Date()
	if year != 2016 || month != time.June || day != 18 {
		t.Fatalf(
			"Entry.PublishDate seems odd: %d/%d/%d, expected: 2016/6/18",
			year, month, day,
		)
	}
}
