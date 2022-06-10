package gui

var successMessage string

func UpdateSuccessMessage(message string) {
	successMessage = message
	NotifyContentChange_Success_text_SucessLabel(message)
}
