package gui

import (
	"github.com/DanielHauge/ge-go-sdl2"
	"github.com/veandco/go-sdl2/sdl"
)

// UI element declaration
var (
	UICompiler_text_label           ge_go_sdl2.Text
	UICompiler_text_label_2         ge_go_sdl2.Text
	UICompiler_container_InputView  ge_go_sdl2.Container
	UICompiler_container_StatusView ge_go_sdl2.Container
	UICompiler_text_label_3         ge_go_sdl2.Text
	View_UICompiler                 ge_go_sdl2.View
)

// UICompiler view and elements construction
func construct_View_UICompiler() {

	// label text construction
	UICompiler_text_label = ge_go_sdl2.Text{
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
	UICompiler_text_label_2 = ge_go_sdl2.Text{
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
	UICompiler_container_InputView = ge_go_sdl2.Container{
		X:      10,
		Y:      110,
		Id:     "InputView",
		ViewId: "Input",
	}

	// StatusView container construction
	UICompiler_container_StatusView = ge_go_sdl2.Container{
		X:      10,
		Y:      280,
		Id:     "StatusView",
		ViewId: "Status Views (Placeholder)",
	}

	// label_3 text construction
	UICompiler_text_label_3 = ge_go_sdl2.Text{
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

	// UICompiler view element construction
	var view_UICompiler_children []interface{}
	view_UICompiler_children = append(view_UICompiler_children, UICompiler_text_label)
	view_UICompiler_children = append(view_UICompiler_children, UICompiler_text_label_2)
	view_UICompiler_children = append(view_UICompiler_children, UICompiler_container_InputView)
	view_UICompiler_children = append(view_UICompiler_children, UICompiler_container_StatusView)
	view_UICompiler_children = append(view_UICompiler_children, UICompiler_text_label_3)
	View_UICompiler = ge_go_sdl2.View{
		X:           0,
		Y:           0,
		H:           463,
		W:           569,
		Id:          "UICompiler",
		BorderColor: DefaultViewBorderColor,
		BgColor:     DefaultViewColor,
		Children:    view_UICompiler_children}

}
