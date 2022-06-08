package uic_go_sdl2

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var fileNameExtractionRegEx *regexp.Regexp

func Compile(inputDir string, outputDir string) {
	os.Mkdir(outputDir+"/gui", 0755)
	files, err := ioutil.ReadDir(inputDir)
	var viewIdentifiers []string
	nilOrPanic(err)
	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".ui") {
			continue
		}
		viewIdentifier := compileView(fmt.Sprintf("%s/%s", inputDir, file.Name()), outputDir)
		viewIdentifiers = append(viewIdentifiers, viewIdentifier)
	}
	mainCode := constructMain(outputDir+string(os.PathSeparator)+"gui", viewIdentifiers)
	createGoFile(fmt.Sprintf("%s/gui/main.go", outputDir), mainCode)
	copyFont(outputDir)
}

func compileView(inputFile string, outputDir string) string {

	// Open file
	fileName := strings.Split(filepath.Base(inputFile), ".")[0]
	file, err := os.Open(inputFile)
	nilOrPanic(err)
	defer file.Close()

	// Read file content
	byteArray, err := ioutil.ReadAll(file)
	nilOrPanic(err)

	// Unmarshal GUI
	var qtUi QtUi
	err = xml.Unmarshal(byteArray, &qtUi)
	nilOrPanic(err)

	// Parse QT qtUi to ge_go_sdl2 GUI
	view := parseFromQtFormat(qtUi)

	// Generate the code for the view and presenter
	identifier, viewCode, uei, scei := constructView(view)
	presenterCode := constructPresenter(uei, scei)

	// Create go files and format them.
	createGoFile(fmt.Sprintf("%s/gui/%s_view.go", outputDir, fileName), viewCode)
	createGoFile(fmt.Sprintf("%s/gui/%s_presenter.go", outputDir, fileName), presenterCode)

	return identifier
}
