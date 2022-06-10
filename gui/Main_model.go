package gui

import "fmt"

func reportFailIfPanic() {
	if e := recover(); e != nil {
		NotifyViewChange_UICompiler_container_StatusView(View_Fail.Id)
		updateErrorMessage(fmt.Sprintf("%s", e))
	}
}

func reportSuccess(message string) {
	NotifyViewChange_UICompiler_container_StatusView(View_Success.Id)
	UpdateSuccessMessage(message)
}
