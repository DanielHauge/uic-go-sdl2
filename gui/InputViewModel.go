package gui

import (
	"fmt"
	uic_go_sdl2 "uic-go-sdl2"
)

type Arguments struct {
	InputDirectory  string
	OutputDirectory string
}

const (
	blank  = 0
	sucess = 1
	fail   = 2
)

var (
	Args   Arguments
	status = blank
)

func Compile() {
	defer func() {
		if e := recover(); e != nil {
			reportError(fmt.Sprintf("%s", e))
		}
	}()
	viewsCompiled := uic_go_sdl2.Compile(Args.InputDirectory, Args.OutputDirectory)
	reportSuccess(fmt.Sprintf("%d views were compiled.", viewsCompiled))
}

func reportError(message string) {
	if status != fail {
		NotifyViewChange_MainWindow_container_StatusView(View_StatusBadView.Id)
		status = fail
	}
	NotifyContentChange_StatusBadView_text_ErrorLabel(message)
	NotifyContentChange_StatusBadView_text_ErrorLabel(message)
}

func reportSuccess(message string) {
	if status != sucess {
		NotifyViewChange_MainWindow_container_StatusView(View_StatusSuccesView.Id)
		status = sucess
	}
	NotifyContentChange_StatusSuccesView_text_SucessLabel(message)
}
