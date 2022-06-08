package gui

import (
	"github.com/DanielHauge/ge-go-sdl2"
)

// UI element declaration
var (
	View_StatusBlankView ge_go_sdl2.View
)

// StatusBlankView view and elements construction
func construct_View_StatusBlankView() {

	// StatusBlankView view element construction
	var view_StatusBlankView_children []interface{}
	View_StatusBlankView = ge_go_sdl2.View{
		X:           0,
		Y:           0,
		H:           144,
		W:           525,
		Id:          "StatusBlankView",
		BorderColor: DefaultViewBorderColor,
		BgColor:     DefaultViewColor,
		Children:    view_StatusBlankView_children}

}
