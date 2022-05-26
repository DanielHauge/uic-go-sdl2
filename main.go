package uic_go_sdl2

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Parsing GUI XML")
	// Extract filepath
	if len(os.Args) < 2 {
		panic("Not enough arguments, expects a file path as argument")
	}
	filePath := os.Args[1]

	// Open file
	file, err := os.Open(filePath)
	NilOrPanic(err)
	defer file.Close()

	// Read file content
	byteArray, err := ioutil.ReadAll(file)
	NilOrPanic(err)

	// Unmarshal GUI
	var gui interface{}
	err = xml.Unmarshal(byteArray, &gui)
	NilOrPanic(err)

	// Generate interface files
	interfaceFileContent := "package gui - blablah"
	encodedContent := []byte(interfaceFileContent)
	ioutil.WriteFile("filename", encodedContent, 0644) // the 0644 is octal representation of the filemode

}

func NilOrPanic(v interface{}) {
	if v != nil {
		panic(v)
	}
}
