package rulesReader

import (
	creater "Linter/packages/creater"
	_ "embed"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

type Rules struct {
	Rules      []Rule      `json:"rules"`
	Exceptions []Exception `json:"exceptions"`
	Blanklines int         `json:"blanklines"`
}
type Rule struct {
	Id            string `json:"id"`
	Space         bool   `json:"spaceAfter"`
	Indent        int    `json:"indentComingLines"`
	AddBlankLines bool   `json:"addBlankLinesBefore"`
}

type Exception struct {
	Exception string `json:"exception"`
	After     bool   `json:"after"`
}

//go:embed rules.json
var embededRules []byte

func ReadJson(file string) Rules {
	acceptedFormats := map[string]bool{
		".yaml": true,
		".json": true,
	}
	if creater.FileExists(file) {
		if acceptedFormats[path.Ext(file)] {
			file, err := os.Open(file)
			if err != nil {
				log.Fatal(err)
			}
			byteValue, e := ioutil.ReadAll(file)
			if e != nil {
				log.Fatal(e)
			}
			var rules Rules
			jsonErr := json.Unmarshal(byteValue, &rules)

			if jsonErr != nil {
				log.Fatal(jsonErr)
			}
			file.Close()
			return rules
		}
		fmt.Println("Not a .yaml or .json file\nTry again with a .json or .yaml file")
		os.Exit(0)
		return Rules{}
	}

	var rules Rules
	err := json.Unmarshal(embededRules, &rules)
	if err != nil {
		log.Fatal(err)
	}
	return rules
}
