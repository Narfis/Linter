package tests

import (
	"Linter/packages/rulesReader"
	"errors"
	"reflect"
	"testing"
)

func TestReadJson(t *testing.T) {

	var tests = []struct {
		input    string
		expected error
	}{
		{"../packages/rulesReader/rules.json", nil},
		{"../rules.notJson", errors.New("")},
		{"../rules.yaml", errors.New("")},
	}

	for _, test := range tests {
		_, err := rulesReader.ReadJson(test.input)
		if reflect.TypeOf(err) != reflect.TypeOf(test.expected) {
			t.Errorf("Not correct, rulesReader gave %s, while we were expecting %s", reflect.TypeOf(err), test.expected)
		}
	}
}
