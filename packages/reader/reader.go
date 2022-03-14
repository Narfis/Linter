package reader

import (
	"Linter/packages/creater"
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
)

//Reads from a file and returns its content to a []string
func ReadFile(file string, acceptedFormats map[string]bool) ([]string, error) {
	if !creater.FileExists(file) {
		return nil, errors.New("input file doesn't exist")
	}
	if acceptedFormats[path.Ext(file)] || acceptedFormats == nil {
		f, e := os.Open(file)

		if e != nil {
			log.Fatal(e)
		}
		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanLines)
		var textlines []string

		for scanner.Scan() {
			textlines = append(textlines, (scanner.Text() + "\n"))
		}
		f.Close()
		return textlines, nil
	}
	return nil, fmt.Errorf("not an accepted extention")
}
