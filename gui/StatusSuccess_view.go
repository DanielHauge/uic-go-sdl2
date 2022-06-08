package gui

import (
	"github.com/DanielHauge/ge-go-sdl2"
	"github.com/veandco/go-sdl2/sdl"
)

// UI element declaration
var (
	StatusSuccesView_text_label       ge_go_sdl2.Text
	StatusSuccesView_text_SucessLabel ge_go_sdl2.Text
	StatusSuccesView_text_label_2     ge_go_sdl2.Text
	View_StatusSuccesView             ge_go_sdl2.View
)

// StatusSuccesView view and elements construction
func construct_View_StatusSuccesView() {

	// label text construction
	StatusSuccesView_text_label = ge_go_sdl2.Text{
		Id:        "label",
		Content:   "Success",
		Font:      DefaultFont,
		Size:      34,
		X:         10,
		Y:         10,
		Alignment: ge_go_sdl2.Center,
		TextColor: sdl.Color{R: 0, G: 255, B: 127, A: 255},
	}

	// SucessLabel text construction
	StatusSuccesView_text_SucessLabel = ge_go_sdl2.Text{
		Id:        "SucessLabel",
		Content:   "Sucess message",
		Font:      DefaultFont,
		Size:      DefaultLabelSize,
		X:         10,
		Y:         50,
		Alignment: ge_go_sdl2.Center,
		TextColor: sdl.Color{R: 0, G: 0, B: 0, A: 0},
	}

	// label_2 text construction
	StatusSuccesView_text_label_2 = ge_go_sdl2.Text{
		Id:        "label_2",
		Content:   "（*＾-＾*）(*^_^*)（*＾-＾*）(*^_^*)（*＾-＾*）(*^_^*)（*＾-＾*）(*^_^*)（*＾-＾*）(*^_^*)",
		Font:      DefaultFont,
		Size:      12,
		X:         10,
		Y:         100,
		Alignment: ge_go_sdl2.Center,
		TextColor: sdl.Color{R: 0, G: 0, B: 0, A: 0},
	}

	// StatusSuccesView view element construction
	var view_StatusSuccesView_children []interface{}
	view_StatusSuccesView_children = append(view_StatusSuccesView_children, StatusSuccesView_text_label)
	view_StatusSuccesView_children = append(view_StatusSuccesView_children, StatusSuccesView_text_SucessLabel)
	view_StatusSuccesView_children = append(view_StatusSuccesView_children, StatusSuccesView_text_label_2)
	View_StatusSuccesView = ge_go_sdl2.View{
		X:           0,
		Y:           0,
		H:           142,
		W:           521,
		Id:          "StatusSuccesView",
		BorderColor: DefaultViewBorderColor,
		BgColor:     DefaultViewColor,
		Children:    view_StatusSuccesView_children}

}
