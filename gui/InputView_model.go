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
	defer reportIfError()
	viewsCompiled := uic_go_sdl2.Compile(CompilerArgs.InputDirectory, CompilerArgs.OutputDirectory)
	reportSuccess(fmt.Sprintf("%d views were compiled.", viewsCompiled))
}

func reportIfError() {
	if e := recover(); e != nil {
		NotifyViewChange_MainWindow_container_StatusView(View_StatusBadView.Id)
		NotifyContentChange_StatusBadView_text_ErrorLabel(fmt.Sprintf("%s", e))
	}
}

func reportSuccess(message string) {
	NotifyViewChange_MainWindow_container_StatusView(View_StatusSuccesView.Id)
	NotifyContentChange_StatusSuccesView_text_SucessLabel(message)
}
