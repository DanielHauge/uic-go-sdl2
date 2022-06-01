package gui

import ge_go_sdl2 "github.com/DanielHauge/ge-go-sdl2"

func OnClick_Hello_button_pushButton() {
	// Update model
}

func OnNewValue_Hello_textField_textEdit(value string) {
	// Update model
}

func NotifyValueChange_Hello_textField_textEdit(value string) {
	ge_go_sdl2.NotifyPropertyChangeAsync(Hello_textField_textEdit.Id, "Value", value)
}
func NotifyContentChange_Hello_text_label(value string) {
	ge_go_sdl2.NotifyPropertyChangeAsync(Hello_text_label.Id, "Content", value)
}
func NotifyViewChange_Hello_container_groupBox(value string) {
	ge_go_sdl2.NotifyPropertyChangeAsync(Hello_container_groupBox.Id, "ViewId", value)
}
