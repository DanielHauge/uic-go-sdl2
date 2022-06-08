package gui

import "github.com/DanielHauge/ge-go-sdl2"

func NotifyContentChange_StatusSuccesView_text_label(value string) {
	ge_go_sdl2.NotifyPropertyChangeAsync(StatusSuccesView_text_label.Id, "Content", value)
}
func NotifyContentChange_StatusSuccesView_text_SucessLabel(value string) {
	ge_go_sdl2.NotifyPropertyChangeAsync(StatusSuccesView_text_SucessLabel.Id, "Content", value)
}
func NotifyContentChange_StatusSuccesView_text_label_2(value string) {
	ge_go_sdl2.NotifyPropertyChangeAsync(StatusSuccesView_text_label_2.Id, "Content", value)
}
