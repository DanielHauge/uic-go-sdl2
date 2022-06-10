package gui

import (
	"github.com/DanielHauge/ge-go-sdl2"
	"github.com/veandco/go-sdl2/sdl"
)

// UI element declaration
var (
	Success_text_label       ge_go_sdl2.Text
	Success_text_SucessLabel ge_go_sdl2.Text
	Success_text_label_2     ge_go_sdl2.Text
	View_Success             ge_go_sdl2.View
)

// Success view and elements construction
func construct_View_Success() {

	// label text construction
	Success_text_label = ge_go_sdl2.Text{
		Id:      "label",
		Content: "Success",
		Font:    DefaultFont,
		Size:    34,
		X:       10,
		Y:       10, W: 501,
		H: 31,

		Alignment: ge_go_sdl2.Center,
		TextColor: sdl.Color{R: 0, G: 255, B: 127, A: 255},
	}

	// SucessLabel text construction
	Success_text_SucessLabel = ge_go_sdl2.Text{
		Id:      "SucessLabel",
		Content: "Sucess message",
		Font:    DefaultFont,
		Size:    DefaultLabelSize,
		X:       10,
		Y:       50, W: 501,
		H: 41,

		Alignment: ge_go_sdl2.Center,
		TextColor: sdl.Color{R: 0, G: 0, B: 0, A: 0},
	}

	// label_2 text construction
	Success_text_label_2 = ge_go_sdl2.Text{
		Id:      "label_2",
		Content: "┏(^0^)┛ - (/^_^)/ - ┗(^0^)┓ - (^_^) - ┏(^0^)┛ - (/^_^)/ - ┗(^0^)┓",
		Font:    DefaultFont,
		Size:    12,
		X:       10,
		Y:       100, W: 501,
		H: 20,

		Alignment: ge_go_sdl2.Center,
		TextColor: sdl.Color{R: 0, G: 0, B: 0, A: 0},
	}

	// Success view element construction
	var view_Success_children []interface{}
	view_Success_children = append(view_Success_children, Success_text_label)
	view_Success_children = append(view_Success_children, Success_text_SucessLabel)
	view_Success_children = append(view_Success_children, Success_text_label_2)
	View_Success = ge_go_sdl2.View{
		X:           0,
		Y:           0,
		H:           142,
		W:           521,
		Id:          "Success",
		BorderColor: DefaultViewBorderColor,
		BgColor:     DefaultViewColor,
		Children:    view_Success_children}

}
