package gui

import (
	"github.com/DanielHauge/ge-go-sdl2"
	"github.com/veandco/go-sdl2/sdl"
)

// UI element declaration
var (
	MainWindow_text_label           ge_go_sdl2.Text
	MainWindow_text_label_2         ge_go_sdl2.Text
	MainWindow_container_InputView  ge_go_sdl2.Container
	MainWindow_container_StatusView ge_go_sdl2.Container
	MainWindow_text_label_3         ge_go_sdl2.Text
	View_MainWindow                 ge_go_sdl2.View
)

// MainWindow view and elements construction
func construct_View_MainWindow() {

	// label text construction
	MainWindow_text_label = ge_go_sdl2.Text{
		Id:      "label",
		Content: "UI Compiler",
		Font:    DefaultFont,
		Size:    24,
		X:       10,
		Y:       10, W: 271,
		H: 31,

		Alignment: ge_go_sdl2.Left,
		TextColor: sdl.Color{R: 0, G: 0, B: 0, A: 0},
	}

	// label_2 text construction
	MainWindow_text_label_2 = ge_go_sdl2.Text{
		Id:      "label_2",
		Content: "Source code available at: https://github.com/DanielHauge/uic-go-sdl2",
		Font:    DefaultFont,
		Size:    14,
		X:       10,
		Y:       40, W: 551,
		H: 31,

		Alignment: ge_go_sdl2.Left,
		TextColor: sdl.Color{R: 0, G: 0, B: 0, A: 0},
	}

	// InputView container construction
	MainWindow_container_InputView = ge_go_sdl2.Container{
		X:      10,
		Y:      110,
		Id:     "InputView",
		ViewId: "InputView",
	}

	// StatusView container construction
	MainWindow_container_StatusView = ge_go_sdl2.Container{
		X:      10,
		Y:      280,
		Id:     "StatusView",
		ViewId: "StatusBlankView",
	}

	// label_3 text construction
	MainWindow_text_label_3 = ge_go_sdl2.Text{
		Id:      "label_3",
		Content: "GUI envionment at: https://github.com/DanielHauge/ge-go-sdl2",
		Font:    DefaultFont,
		Size:    14,
		X:       10,
		Y:       70, W: 551,
		H: 31,

		Alignment: ge_go_sdl2.Left,
		TextColor: sdl.Color{R: 0, G: 0, B: 0, A: 0},
	}

	// MainWindow view element construction
	var view_MainWindow_children []interface{}
	view_MainWindow_children = append(view_MainWindow_children, MainWindow_text_label)
	view_MainWindow_children = append(view_MainWindow_children, MainWindow_text_label_2)
	view_MainWindow_children = append(view_MainWindow_children, MainWindow_container_InputView)
	view_MainWindow_children = append(view_MainWindow_children, MainWindow_container_StatusView)
	view_MainWindow_children = append(view_MainWindow_children, MainWindow_text_label_3)
	View_MainWindow = ge_go_sdl2.View{
		X:           0,
		Y:           0,
		H:           463,
		W:           569,
		Id:          "MainWindow",
		BorderColor: DefaultViewBorderColor,
		BgColor:     DefaultViewColor,
		Children:    view_MainWindow_children}

}
