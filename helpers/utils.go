package helpers

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func GetInput(fp string) string {
	content, err := ioutil.ReadFile(fp)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
