package tests

import (
	"Linter/packages/creater"
	"Linter/packages/reader"
	"Linter/packages/rulesReader"
	"Linter/packages/theLinter"
	"Linter/packages/writer"
	"strings"
	"testing"
)

func TestLinter(t *testing.T) {
	acceptedFormats := map[string]bool{
		".tex":  true,
		".bibz": true,
		".tikz": true,
	}
	rules, err := rulesReader.ReadJson("../packages/rulesReader/rules.json")
	var tests = []struct {
		input    string
		expected string
	}{
		{"\\begin\nhello", "\\begin\n\thello"},
		{"%hello", "% hello"},
		{"\\%hello", "\\%hello"},
		{"\\section", strings.Repeat("\n", rules.Blanklines) + "\\section"},
		{"\\begin{document}\nhello\n\\begin{article}\nshould indent\n\\end\n\\end{document}", "\\begin{document}\nhello\n\\begin{article}\n\tshould indent\n\\end\n\\end{document}"},
		{"        this should fix itself", "this should fix itself"},
		{"This should fix spacing\n\n\n\n\nhello", "This should fix spacing\n\nhello"},
		{"hello%\\begin\nhello", "hello % \\begin\nhello"},
		{"%\\begin\nhello", "% \\begin\nhello"},
	}

	if err != nil {
		t.Error("Failed on the wrong thing, rules didn't load")
	}
	headers := false
	for _, test := range tests {
		creater.CreateFile("test.tex", acceptedFormats)
		file := writer.OpenFile("test.tex")
		writer.WriteTo(test.input, file)
		theLinter.DoLint("test.tex", "out.tex", rules, headers)

		fileRead, err := reader.ReadFile("out.tex", acceptedFormats)
		if err != nil {
			t.Error("Error in readfile")
		}
		wholeString := ""
		for _, word := range fileRead {
			wholeString += word
		}
		if wholeString != test.expected+"\n" {
			t.Error("Not correctly linted")
		}

	}
}
