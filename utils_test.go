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

func TestRemoveHash(t *testing.T) {
	urlOrig := "http://google.com"
	urlResult := RemoveHash(urlOrig)
	if urlResult != "http://google.com" {
		t.Error("url is invalid.", urlOrig, urlResult)
	}

	urlOrig = "http://google.com#hash"
	urlResult = RemoveHash(urlOrig)
	if urlResult != "http://google.com" {
		t.Error("url is invalid.", urlOrig, urlResult)
	}
}

func TestRemoveTwitterUrlHash(t *testing.T) {
	urlOrig := "http://google.com"
	urlResult := RemoveTwitterUrlHash(urlOrig)
	if urlResult != "http://google.com" {
		t.Error("url is invalid.", urlOrig, urlResult)
	}

	urlOrig = "http://google.com#hash"
	urlResult = RemoveTwitterUrlHash(urlOrig)
	if urlResult != "http://google.com#hash" {
		t.Error("url is invalid.", urlOrig, urlResult)
	}

	urlOrig = "http://google.com#.XXXX.twitter"
	urlResult = RemoveTwitterUrlHash(urlOrig)
	if urlResult != "http://google.com" {
		t.Error("url is invalid.", urlOrig, urlResult)
	}
}

func TestRemoveUtmParams(t *testing.T) {
	urlOrig := "http://google.com"
	urlResult := RemoveUtmParams(urlOrig)
	if urlResult != "http://google.com" {
		t.Error("url is invalid.", urlOrig, urlResult)
	}

	urlOrig = "http://google.com?key=val"
	urlResult = RemoveUtmParams(urlOrig)
	if urlResult != "http://google.com?key=val" {
		t.Error("url is invalid.", urlOrig, urlResult)
	}

	urlOrig = "http://google.com?utm_content=aaa"
	urlResult = RemoveUtmParams(urlOrig)
	if urlResult != "http://google.com" {
		t.Error("url is invalid.", urlOrig, urlResult)
	}

	urlOrig = "http://google.com?key=val&key2=val2"
	urlResult = RemoveUtmParams(urlOrig)
	if urlResult != "http://google.com?key=val&key2=val2" {
		t.Error("url is invalid.", urlOrig, urlResult)
	}

	urlOrig = "http://google.com?key=val&utm_content=aaa"
	urlResult = RemoveUtmParams(urlOrig)
	if urlResult != "http://google.com?key=val" {
		t.Error("url is invalid.", urlOrig, urlResult)
	}

	urlOrig = "http://google.com?utm_content=aaa&key=val"
	urlResult = RemoveUtmParams(urlOrig)
	if urlResult != "http://google.com?key=val" {
		t.Error("url is invalid.", urlOrig, urlResult)
	}

	urlOrig = "http://google.com?utm_content=aaa&utm_medium=bbb"
	urlResult = RemoveUtmParams(urlOrig)
	if urlResult != "http://google.com" {
		t.Error("url is invalid.", urlOrig, urlResult)
	}
}
