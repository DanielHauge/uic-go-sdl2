package gui

import "github.com/DanielHauge/ge-go-sdl2"

func NotifyContentChange_MainWindow_text_label(value string) {
	ge_go_sdl2.NotifyPropertyChangeAsync(MainWindow_text_label.Id, "Content", value)
}
func NotifyContentChange_MainWindow_text_label_2(value string) {
	ge_go_sdl2.NotifyPropertyChangeAsync(MainWindow_text_label_2.Id, "Content", value)
}
func NotifyViewChange_MainWindow_container_InputView(value string) {
	ge_go_sdl2.NotifyPropertyChangeAsync(MainWindow_container_InputView.Id, "ViewId", value)
}
func NotifyViewChange_MainWindow_container_StatusView(value string) {
	ge_go_sdl2.NotifyPropertyChangeAsync(MainWindow_container_StatusView.Id, "ViewId", value)
}
func NotifyContentChange_MainWindow_text_label_3(value string) {
	ge_go_sdl2.NotifyPropertyChangeAsync(MainWindow_text_label_3.Id, "Content", value)
}
