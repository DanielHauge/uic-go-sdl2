package gui

import (
	"fmt"
	uic_go_sdl2 "uic-go-sdl2"
)

var CompilerArgs struct {
	InputDirectory  string
	OutputDirectory string
}

func Compile() {
	defer reportFailIfPanic() // From Main_model.go
	viewsCompiled := uic_go_sdl2.Compile(CompilerArgs.InputDirectory, CompilerArgs.OutputDirectory)
	reportSuccess(fmt.Sprintf("%d views were compiled.", viewsCompiled)) // From Main_model.go
}
