package tests

import (
	"Linter/packages/reader"
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
		{"testingFiles/bigtext.tex", nil},
		{"testingFiles/emptyDot.tex", nil},
		{"testingFiles/hello.tex", nil},
		{"testingFiles/macketest.tex", nil},
		{"testingFiles/sectionTest.tex", nil},
		{"testingFiles/test.tex", nil},
		{"testingFiles/try.tex", nil},
		{"testingFiles/bigtext.tex", nil},
		{"testingFiles/notTex.txt", errors.New("")},
		{"testingFiles/notAnActualFile.bibz", errors.New("")},
		{"oog.tex", errors.New("")},
		{"", errors.New("")},
	}
	for _, test := range tests {
		_, err := reader.ReadFile(test.input, acceptedFormats)
		if reflect.TypeOf(err) != reflect.TypeOf(test.expected) {
			t.Errorf("Not correct, reader gave the error : %s, at the file : \"%s\"", err, test.input)
		}
	}

}
