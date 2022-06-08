package gui

import "github.com/DanielHauge/ge-go-sdl2"

func NotifyContentChange_StatusBadView_text_label(value string) {
	ge_go_sdl2.NotifyPropertyChangeAsync(StatusBadView_text_label.Id, "Content", value)
}
func NotifyContentChange_StatusBadView_text_ErrorLabel(value string) {
	ge_go_sdl2.NotifyPropertyChangeAsync(StatusBadView_text_ErrorLabel.Id, "Content", value)
}
func NotifyContentChange_StatusBadView_text_label_2(value string) {
	ge_go_sdl2.NotifyPropertyChangeAsync(StatusBadView_text_label_2.Id, "Content", value)
}
