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
