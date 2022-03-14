package creater

import (
	"fmt"
	"os"
	"path"
)

//Creates a file within the acceptedFormats.
func CreateFile(file string, acceptedFormats map[string]bool) error {
	if acceptedFormats[path.Ext(file)] {
		os.Create(file)
		return nil
	}
	return fmt.Errorf("file you're trying to create is not within acceptedFormats")
}

//Checks if file exists.
func FileExists(file string) bool {
	info, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
