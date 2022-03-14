package writer

import (
	"io"
	"log"
	"os"
)

//Opens a file that can be written to.
func OpenFile(toFile string) *os.File {
	file, err := os.OpenFile(toFile, os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

//Writes to file
func WriteTo(text string, f *os.File) {
	_, err := io.WriteString(f, text)

	if err != nil {
		log.Fatal(err)
	}
}
