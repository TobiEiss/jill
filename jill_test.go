package jill_test

import (
	"testing"

	"github.com/TobiEiss/jill"
)

func TestJill(t *testing.T) {
	container, err := jill.ParseJSON([]byte(`{
		"outter":{
			"inner":{
				"value1":10,
				"value2":22
			},
			"alsoInner":{
				"value1":20
			}
		}
	}`))

	// check error
	if err != nil {
		t.Errorf("Error while parse json: %s", err)
	}

	// query
	result, err := container.Query("ADD ( outter.inner.value1, outter.alsoInner.value1 )")
	if err != nil {
		t.Errorf("Error while query json: %s", err)
	}
	if result.(float64) != 30 {
		t.Errorf("Not expected result")
	}
}
