package creater

import (
	"fmt"
	"os"
	"path"
)

func CreateFile(file string, acceptedFormats map[string]bool) {
	if acceptedFormats[path.Ext(file)] {
		os.Create(file)
		return
	}
	fmt.Println("The output isn't a .tex file, exiting")
	os.Exit(0)
}

func FileExists(file string) bool {
	info, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
