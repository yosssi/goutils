package goutils

import (
	"reflect"
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
