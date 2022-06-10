package gui

import (
	"github.com/DanielHauge/ge-go-sdl2"
	"github.com/veandco/go-sdl2/sdl"
)

// UI element declaration
var (
	Fail_text_label      ge_go_sdl2.Text
	Fail_text_ErrorLabel ge_go_sdl2.Text
	Fail_text_label_2    ge_go_sdl2.Text
	View_Fail            ge_go_sdl2.View
)

// Fail view and elements construction
func construct_View_Fail() {

	// label text construction
	Fail_text_label = ge_go_sdl2.Text{
		Id:      "label",
		Content: "Ups... Something went wrong",
		Font:    DefaultFont,
		Size:    21,
		X:       10,
		Y:       10, W: 311,
		H: 31,

		Alignment: ge_go_sdl2.Left,
		TextColor: sdl.Color{R: 255, G: 85, B: 0, A: 255},
	}

	// ErrorLabel text construction
	Fail_text_ErrorLabel = ge_go_sdl2.Text{
		Id:      "ErrorLabel",
		Content: "Error message",
		Font:    DefaultFont,
		Size:    DefaultLabelSize,
		X:       10,
		Y:       80, W: 471,
		H: 16,

		Alignment: ge_go_sdl2.Left,
		TextColor: sdl.Color{R: 0, G: 0, B: 0, A: 0},
	}

	// label_2 text construction
	Fail_text_label_2 = ge_go_sdl2.Text{
		Id:      "label_2",
		Content: "(\\*.*)\\┳━┳        ->       (╯°□°)╯^┻━┻     ->    ╰(°□°)╯",
		Font:    DefaultFont,
		Size:    DefaultLabelSize,
		X:       10,
		Y:       50, W: 501,
		H: 20,

		Alignment: ge_go_sdl2.Left,
		TextColor: sdl.Color{R: 0, G: 0, B: 0, A: 0},
	}

	// Fail view element construction
	var view_Fail_children []interface{}
	view_Fail_children = append(view_Fail_children, Fail_text_label)
	view_Fail_children = append(view_Fail_children, Fail_text_ErrorLabel)
	view_Fail_children = append(view_Fail_children, Fail_text_label_2)
	View_Fail = ge_go_sdl2.View{
		X:           0,
		Y:           0,
		H:           143,
		W:           523,
		Id:          "Fail",
		BorderColor: DefaultViewBorderColor,
		BgColor:     DefaultViewColor,
		Children:    view_Fail_children}

}
