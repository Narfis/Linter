package writer

import (
	"io"
	"log"
	"os"
)

func OpenFile(toFile string) *os.File {
	file, err := os.OpenFile(toFile, os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func WriteToNew(text string, f *os.File) {

	_, err2 := io.WriteString(f, text)

	if err2 != nil {
		log.Fatal(err2)
	}
}