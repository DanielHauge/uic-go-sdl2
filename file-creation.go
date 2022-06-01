package uic_go_sdl2

import (
	"io/ioutil"
	"log"
	"os/exec"
)

func createGoFile(filePath string, content string) {
	encodedContent := []byte(content)
	ioutil.WriteFile(filePath, encodedContent, 0644) // the 0644 is octal representation of the filemode
	cmd := exec.Command("go", "fmt", filePath)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func copyFont(outputDir string) {
	file, err := ioutil.ReadFile("./assets/default.ttf")
	nilOrPanic(err)

	err = ioutil.WriteFile(outputDir+"/default.ttf", file, 0644)
	nilOrPanic(err)
}
