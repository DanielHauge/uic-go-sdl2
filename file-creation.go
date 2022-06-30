package uic_go_sdl2

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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
	path, err := os.Executable()
	dir, _ := filepath.Split(path)
	file, err := ioutil.ReadFile(dir + "default.ttf")
	nilOrPanic(err)

	err = ioutil.WriteFile(outputDir+"/gui/default.ttf", file, 0644)
	nilOrPanic(err)
}
