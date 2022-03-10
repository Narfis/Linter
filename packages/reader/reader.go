package reader

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
)

func ReadFile(text string, acceptedFormats map[string]bool) []string {

	if acceptedFormats[path.Ext(text)] || acceptedFormats == nil {
		f, e := os.Open(text)

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
		return textlines
	}
	fmt.Println("Not a valid file to read from")
	fmt.Println("Try again with a format within the acceptedFormats map")
	os.Exit(0)
	return nil
}
