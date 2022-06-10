package gui

import "github.com/DanielHauge/ge-go-sdl2"

func OnNewValue_Input_textField_textEdit(value string) {
	CompilerArgs.InputDirectory = value
}

func OnNewValue_Input_textField_textEdit_2(value string) {
	CompilerArgs.OutputDirectory = value
}

func OnClick_Input_button_pushButton() {
	Compile()
}

func NotifyValueChange_Input_textField_textEdit(value string) {
	ge_go_sdl2.NotifyPropertyChangeAsync(Input_textField_textEdit.Id, "Value", value)
}
func NotifyContentChange_Input_text_label(value string) {
	ge_go_sdl2.NotifyPropertyChangeAsync(Input_text_label.Id, "Content", value)
}
func NotifyValueChange_Input_textField_textEdit_2(value string) {
	ge_go_sdl2.NotifyPropertyChangeAsync(Input_textField_textEdit_2.Id, "Value", value)
}
func NotifyContentChange_Input_text_label_2(value string) {
	ge_go_sdl2.NotifyPropertyChangeAsync(Input_text_label_2.Id, "Content", value)
}
