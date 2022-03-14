package main

import (
	"Linter/packages/rulesReader"
	linter "Linter/packages/theLinter"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("Linter", "Lints your .tex document")
	file := parser.String("f", "file", &argparse.Options{Required: true, Help: "File to Lint"})
	newFile := parser.String("n", "newFile", &argparse.Options{Required: true, Help: "Output file that will be linted (must end in .tex)"})
	rules := parser.String("r", "rules", &argparse.Options{Required: false, Help: "Own rules (optional)"})
	headers := parser.Flag("", "headers", &argparse.Options{Required: false, Help: "Writes out headers to console"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Println(parser.Usage(err))
		return
	}
	theRules, err := rulesReader.ReadJson(*rules)
	if err != nil {
		log.Fatal(err)
	}
	startTime := time.Now()
	linter.DoLint(*file, *newFile, theRules, *headers)
	duration := time.Since(startTime)
	fmt.Println("The program finished in:", duration.Seconds(), "seconds")
}
