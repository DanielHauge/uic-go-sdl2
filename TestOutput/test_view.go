package gui

import ge_go_sdl2 "github.com/DanielHauge/ge-go-sdl2"

// UI element declaration
var (
	View_ ge_go_sdl2.View
)

//  view and elements construction
func construct__view() {

	//  view element construction
	var view__children []interface{}
	View_ = ge_go_sdl2.View{
		X:           0,
		Y:           0,
		H:           0,
		W:           0,
		Id:          "",
		BorderColor: DefaultViewBorderColor,
		BgColor:     DefaultViewColor,
		Children:    view__children}

}
