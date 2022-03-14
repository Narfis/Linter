package main

import (
	"Linter/packages/reader"
	"Linter/packages/rulesReader"
	"errors"
	"reflect"
	"testing"
)

func TestRead(t *testing.T) {
	acceptedFormats := map[string]bool{
		".tex":  true,
		".bibz": true,
		".tikz": true,
	}
	var tests = []struct {
		input    string
		expected error
	}{
		{"testing/bigtext.tex", nil},
		{"testing/emptyDot.tex", nil},
		{"testing/hello.tex", nil},
		{"testing/macketest.tex", nil},
		{"testing/sectionTest.tex", nil},
		{"testing/test.tex", nil},
		{"testing/try.tex", nil},
		{"testing/bigtext.tex", nil},
		{"testing/notTex.txt", errors.New("")},
		{"testing/notAnActualFile.bibz", errors.New("")},
		{"oog.tex", errors.New("")},
		{"", errors.New("")},
	}
	for _, test := range tests {
		_, err := reader.ReadFile(test.input, acceptedFormats)
		if reflect.TypeOf(err) != reflect.TypeOf(test.expected) {
			t.Errorf("Not correct, reader gave %s, while we were expecting %s", reflect.TypeOf(err), test.expected)
		}
	}

}

func TestReadJson(t *testing.T) {
	var tests = []struct {
		input    string
		expected error
	}{
		{"packages/rulesReader/rules.json", nil},
		{"rules.notJson", errors.New("")},
		{"rules.yaml", errors.New("")},
	}

	for _, test := range tests {
		_, err := rulesReader.ReadJson(test.input)
		if reflect.TypeOf(err) != reflect.TypeOf(test.expected) {
			t.Errorf("Not correct, reader gave %s, while we were expecting %s", reflect.TypeOf(err), test.expected)
		}
	}

}
