package gui

import "github.com/DanielHauge/ge-go-sdl2"

func OnNewValue_InputView_textField_textEdit(value string) {
	Args.InputDirectory = value
}

func OnNewValue_InputView_textField_textEdit_2(value string) {
	Args.OutputDirectory = value
}

func OnClick_InputView_button_pushButton() {
	Compile()
}

func NotifyValueChange_InputView_textField_textEdit(value string) {
	ge_go_sdl2.NotifyPropertyChangeAsync(InputView_textField_textEdit.Id, "Value", value)
}
func NotifyContentChange_InputView_text_label(value string) {
	ge_go_sdl2.NotifyPropertyChangeAsync(InputView_text_label.Id, "Content", value)
}
func NotifyValueChange_InputView_textField_textEdit_2(value string) {
	ge_go_sdl2.NotifyPropertyChangeAsync(InputView_textField_textEdit_2.Id, "Value", value)
}
func NotifyContentChange_InputView_text_label_2(value string) {
	ge_go_sdl2.NotifyPropertyChangeAsync(InputView_text_label_2.Id, "Content", value)
}
