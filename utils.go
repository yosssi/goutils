package goutils

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

const (
	urlPrefixHttp  string = "http://"
	urlPrefixHttps string = "https://"
)

func StructToMap(data interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	elem := reflect.ValueOf(data).Elem()
	size := elem.NumField()

	for i := 0; i < size; i++ {
		field := elem.Type().Field(i).Name
		value := elem.Field(i).Interface()
		m[field] = value
	}

	return m
}

func GetUrls(s string) []string {
	urls := make([]string, 0)
	tokens := strings.Split(s, " ")
	for _, t := range tokens {
		if strings.HasPrefix(t, urlPrefixHttp) || strings.HasPrefix(t, urlPrefixHttps) {
			urls = append(urls, t)
		}
	}
	return urls
}

func NormalUrl(s string) string {
	res, err := http.Get(s)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return res.Request.URL.String()
}

func RemoveHash(s string) string {
	return strings.Split(s, "#")[0]
}

func RemoveTwitterUrlHash(s string) string {
	tokens := strings.Split(s, "#")
	if len(tokens) > 1 && strings.HasPrefix(tokens[1], ".") && strings.HasSuffix(tokens[1], ".twitter") {
		return tokens[0]
	} else {
		return s
	}
}

func RemoveUtmParams(s string) string {
	urlTokens := strings.Split(s, "?")
	params := make([]string, 0)
	if len(urlTokens) > 1 {
		for _, param := range strings.Split(urlTokens[1], "&") {
			if !strings.HasPrefix(param, "utm") {
				params = append(params, param)
			}
		}
	}

	url := urlTokens[0]
	if len(params) > 0 {
		url += "?" + strings.Join(params, "&")
	}
	return url
}
