package goutils

import (
	"fmt"
	"github.com/yosssi/gocmd"
	"os/exec"
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
	output, err := gocmd.Pipe(exec.Command("curl", "-sLI", s), exec.Command("grep", "-E", "Location:|location:"), exec.Command("tail", "-1"))
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	result := string(output)
	if result == "" {
		result = s
	} else {
		result = strings.TrimSpace(strings.TrimPrefix(strings.TrimPrefix(result, "location: "), "Location: "))
	}
	return result
}
