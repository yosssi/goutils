package goutils

import (
	"testing"
)

type TestStruct struct {
	Name    string
	Amount  int
	IsAdmin bool
}

func TestStructToMap(t *testing.T) {
	name := "Taro"
	amount := 5
	isAdmin := true
	testStruct := TestStruct{Name: name, Amount: amount, IsAdmin: isAdmin}
	m := StructToMap(&testStruct)
	if m["Name"] != name || m["Amount"] != amount || m["IsAdmin"] != isAdmin {
		t.Error("The result map is invalid.")
	}
}

func TestGetUrls(t *testing.T) {
	s := ""
	urls := GetUrls(s)
	if len(urls) != 0 {
		t.Error("urls is invalid.")
	}

	s = "example"
	urls = GetUrls(s)
	if len(urls) != 0 {
		t.Error("urls is invalid.")
	}

	s = "http://example.com"
	urls = GetUrls(s)
	if len(urls) != 1 || urls[0] != "http://example.com" {
		t.Error("urls is invalid.")
	}

	s = "https://example.com"
	urls = GetUrls(s)
	if len(urls) != 1 || urls[0] != "https://example.com" {
		t.Error("urls is invalid.")
	}

	s = "example example"
	urls = GetUrls(s)
	if len(urls) != 0 {
		t.Error("urls is invalid.")
	}

	s = "example http://example.com"
	urls = GetUrls(s)
	if len(urls) != 1 || urls[0] != "http://example.com" {
		t.Error("urls is invalid.")
	}

	s = "example https://example.com"
	urls = GetUrls(s)
	if len(urls) != 1 || urls[0] != "https://example.com" {
		t.Error("urls is invalid.")
	}

	s = "http://example.com http://example.com"
	urls = GetUrls(s)
	if len(urls) != 2 || urls[0] != "http://example.com" || urls[1] != "http://example.com" {
		t.Error("urls is invalid.")
	}

	s = "http://example.com https://example.com"
	urls = GetUrls(s)
	if len(urls) != 2 || urls[0] != "http://example.com" || urls[1] != "https://example.com" {
		t.Error("urls is invalid.")
	}

	s = "https://example.com https://example.com"
	urls = GetUrls(s)
	if len(urls) != 2 || urls[0] != "https://example.com" || urls[1] != "https://example.com" {
		t.Error("urls is invalid.")
	}
}

func TestNormalUrl(t *testing.T) {
	url := NormalUrl("http://t.co/cF6LQoH25w")
	if url != "http://play.golang.org/p/0jVsXA2_R0" {
		t.Error("urls is invalid.")
	}

	url = NormalUrl("http://www.google.co.jp/")
	if url != "http://www.google.co.jp/" {
		t.Error("urls is invalid.")
	}
}
