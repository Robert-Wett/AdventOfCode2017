package helpers

import (
	"io/ioutil"
	"log"
)

func GetInput(fp string) string {
	content, err := ioutil.ReadFile(fp)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
