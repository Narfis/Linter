package theLinter

import (
	file "Linter/packages/creater"
	"Linter/packages/reader"
	"Linter/packages/rulesReader"
	"Linter/packages/writer"
	"fmt"
	"log"
	"os"
	"strings"
)

func ModifyOutput(rules rulesReader.Rules, text []string, file *os.File, headers bool) {
	var indents int
	var removeEmpty int

	headerIndents := false
	headerSpaces := false
	headerBlankLines := false

	for _, lines := range text {
		if indents > 0 {
			headerIndents = true
		}
		lines = strings.TrimLeft(lines, " ")
		lines = strings.TrimLeft(lines, "\t")
		lines = AddIndents(lines, indents)
		lines = AddBlankLines(lines, &removeEmpty)
		for i := 0; i < len(rules.Rules); i++ {
			if strings.Contains(lines, rules.Rules[i].Id) {
				lines = DoTheRules(lines, rules, rules.Rules[i], &indents, &headerSpaces, &headerBlankLines)
			}
		}
		lines = NewlineAfterDot(lines, indents, ". ", ".\n")
		if removeEmpty <= 1 {
			writer.WriteToNew(lines, file)
		}
	}

	headerTracker := []bool{headerIndents, headerSpaces, headerBlankLines}
	if headers {
		WriteOutHeaders(headerTracker)
	}
}

func WriteOutHeaders(headerTracker []bool) {
	if headerTracker[0] {
		fmt.Println("* Indents are added")
	}
	if headerTracker[1] {
		fmt.Println("* Spaces are added")

	}
	if headerTracker[2] {
		fmt.Println("* Blanklines are added")
	}
}

func AddLinesBefore(line string, blanks int) string {
	line = strings.Repeat("\n", blanks) + line
	return line
}

func DoTheRules(line string, rules rulesReader.Rules, rule rulesReader.Rule, indents *int, headerSpaces *bool, headerBlankLines *bool) string {
	exception := FindExceptions(line, rules, rule.Id)
	realComment := IndexRealComment(line)
	if realComment == 0 {
		realComment = len(line)
	}
	lineToLint := line[:realComment]
	endofLine := line[realComment:]
	if exception == "" {
		if strings.Contains(lineToLint, rule.Id) {
			*indents += rule.Indent
			if *indents < 0 {
				*indents = 0
			}
			if rule.Space {
				lineToLint = AddSpace(lineToLint, rule, endofLine)
				*headerSpaces = true
			}
			if rule.AddBlankLines {
				lineToLint = AddLinesBefore(lineToLint, rules.Blanklines)
				*headerBlankLines = true
			}
			if rule.Indent < 0 {
				lineToLint = strings.Replace(lineToLint, "\t", "", 1)
			}

		}
	} else {
		//Seperate the line by exceptions
		var wholeLine []string
		copyLine := lineToLint
		for notDone := true; notDone; notDone = (exception != "") {
			indexException := strings.Index(copyLine, exception)
			wholeLine = append(wholeLine, copyLine[:indexException], exception)

			copyLine = copyLine[indexException+len(exception):]
			exception = FindExceptions(copyLine, rules, rule.Id)
		}
		wholeLine = append(wholeLine, copyLine)
		newLine := ""
		for _, word := range wholeLine {
			if FindExceptions(word, rules, rule.Id) == "" {
				if strings.Contains(word, rule.Id) {
					*indents += rule.Indent
					if *indents < 0 {
						*indents = 0
					}
					if rule.Space {
						word = AddSpace(word, rule, lineToLint)
						*headerSpaces = true

					}
					if rule.AddBlankLines {
						lineToLint = AddLinesBefore(lineToLint, rules.Blanklines)
						*headerBlankLines = true

					}
					if rule.Indent < 0 {
						word = strings.Replace(word, "\t", "", 1)
					}
				}
			}
			newLine += word
		}
		lineToLint = newLine

	}
	if strings.Contains(endofLine, rule.Id) {
		if rule.Space {
			endofLine = AddSpace(endofLine, rule, lineToLint)
			*headerSpaces = true
		}
	}

	lineToLint += endofLine
	return lineToLint
}

func FindExceptions(lines string, exception rulesReader.Rules, ruleID string) string {
	for i := 0; i < len(exception.Exceptions); i++ {
		if strings.Contains(lines, ruleID+exception.Exceptions[i].Exception) && exception.Exceptions[i].After {
			return ruleID + exception.Exceptions[i].Exception
		}
		if strings.Contains(lines, (exception.Exceptions[i].Exception+ruleID)) && !exception.Exceptions[i].After {
			return exception.Exceptions[i].Exception + ruleID
		}
	}
	return ""
}

func RemoveDuplicateValues(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}

	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func IndexRealComment(line string) int {
	commentCount := strings.Count(line, "%")

	totalIndex := 0
	copyline := line
	for i := 0; i < commentCount; i++ {
		commentIndex := strings.Index(copyline, "%")
		fakeIndex := strings.Index(copyline, "\\%")
		if fakeIndex+1 == commentIndex {
			copyline = copyline[commentIndex+1:]
			totalIndex += commentIndex + 1
		} else {
			totalIndex += commentIndex
			return totalIndex
		}
	}
	return 0
}

func NewlineAfterDot(lines string, indents int, stringToReplace string, stringReplacer string) string {
	commentCount := strings.Count(lines, "%")
	fakeCount := strings.Count(lines, "\\%")
	stringReplacer += strings.Repeat("\t", indents)
	lines = RemoveSpaces(lines, stringToReplace)
	if commentCount > fakeCount {

		totalIndex := IndexRealComment(lines)
		copyTotal := totalIndex - 1
		if copyTotal > 0 {
			currVal := lines[copyTotal]
			for notDone := true; notDone; notDone = currVal == ' ' {
				copyTotal -= 1
				currVal = lines[copyTotal]
				if currVal != ' ' {
					if currVal == '.' {
						copyTotal = 1
						break
					}
					copyTotal = 0
					break
				}

			}
		} else {
			copyTotal = 0
		}

		countReplaces := strings.Count(lines[:totalIndex], stringToReplace) - copyTotal
		lines = strings.Replace(lines[:totalIndex], stringToReplace, stringReplacer, countReplaces) + lines[totalIndex:]
	} else {
		lines = strings.ReplaceAll(lines, stringToReplace, stringReplacer)
	}

	return lines
}

func CountLeadingSpaces(line string) int {
	return len(line) - len(strings.TrimLeft(line, " "))
}

func RemoveSpaces(lines string, toFind string) string {
	var indexFound []int
	copy := lines

	for i := 0; i < len(lines); i++ {
		found := strings.Index(copy, toFind)
		if found == -1 {
			break
		}
		byteCopy := []byte(copy)
		byteCopy[i] = ' '
		copy = string(byteCopy)
		indexFound = append(indexFound, found)
	}
	indexFound = RemoveDuplicateValues(indexFound)
	totalCompressed := 0
	for i := 0; i < len(indexFound); i++ {
		var compressedLine string

		compressedLine = lines[indexFound[i]-totalCompressed+len(toFind):]

		compressedLen := CountLeadingSpaces(compressedLine)
		compressedLine = strings.TrimLeft(compressedLine, " ")

		lines = lines[:indexFound[i]-totalCompressed+len(toFind)] + compressedLine

		totalCompressed += compressedLen
	}
	return lines
}

func AddIndents(lines string, indents int) string {
	index := strings.Count(lines, "\n") - 1
	lines = lines[:index] + strings.Repeat("\t", indents) + lines[index:]
	return lines
}

func AddSpace(lines string, rules rulesReader.Rule, beforeComment string) string {
	index := strings.Index(lines, "%")
	before := ""
	after := ""
	if len(beforeComment) > 0 {
		if beforeComment[len(beforeComment)-1] != ' ' {
			before = " "
		}
	}

	if len(lines) > index && lines[index+1] != ' ' {
		after = " "
	}
	lines = strings.Replace(lines, rules.Id, (before + rules.Id + after), 1)

	return lines
}

func AddBlankLines(lines string, spaces *int) string {
	if lines[0] == '\n' {
		*spaces += 1
	} else {
		*spaces = 0
	}

	return lines
}

func DoLint(readFrom string, writeTo string, rules rulesReader.Rules, headers bool) {
	acceptedFormats := map[string]bool{
		".tex":  true,
		".bibz": true,
		".tikz": true,
	}
	theFile, err := reader.ReadFile(readFrom, acceptedFormats)
	if err != nil {
		log.Fatal(err)
	}
	err = file.CreateFile(writeTo, acceptedFormats)
	if err != nil {
		log.Fatal(err)
	}
	openFile := writer.OpenFile(writeTo)

	ModifyOutput(rules, theFile, openFile, headers)

	openFile.Close()
}
