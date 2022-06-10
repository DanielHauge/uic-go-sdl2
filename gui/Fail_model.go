package gui

var ErrorMessage string

func updateErrorMessage(message string) {
	ErrorMessage = message
	NotifyContentChange_Fail_text_ErrorLabel(message)
}
