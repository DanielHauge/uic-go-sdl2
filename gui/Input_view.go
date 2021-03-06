package gui

import (
	"github.com/DanielHauge/ge-go-sdl2"
	"github.com/veandco/go-sdl2/sdl"
)

// UI element declaration
var (
	Input_textField_textEdit    ge_go_sdl2.TextField
	Input_text_label            ge_go_sdl2.Text
	Input_textField_textEdit_2  ge_go_sdl2.TextField
	Input_text_label_2          ge_go_sdl2.Text
	Input_text_pushButton_label ge_go_sdl2.Text
	Input_button_pushButton     ge_go_sdl2.Button
	View_Input                  ge_go_sdl2.View
)

// Input view and elements construction
func construct_View_Input() {
	// textEdit text field construction

	textfield_textEdit_onchange := make(chan string)
	Input_textField_textEdit = ge_go_sdl2.TextField{
		X:           10,
		Y:           40,
		H:           31,
		W:           241,
		Id:          "textEdit",
		Size:        10,
		Value:       "",
		BgColor:     DefaultTextFieldColor,
		BorderColor: DefaultTextFieldBorderColor,
		Font:        DefaultFont,
		OnChanged:   textfield_textEdit_onchange,
	}
	ge_go_sdl2.HandleAsCallbackArg(textfield_textEdit_onchange, OnNewValue_Input_textField_textEdit)

	// label text construction
	Input_text_label = ge_go_sdl2.Text{
		Id:      "label",
		Content: "Input directory path",
		Font:    DefaultFont,
		Size:    13,
		X:       20,
		Y:       10, W: 231,
		H: 21,

		Alignment: ge_go_sdl2.Left,
		TextColor: sdl.Color{R: 0, G: 0, B: 0, A: 0},
	}
	// textEdit_2 text field construction

	textfield_textEdit_2_onchange := make(chan string)
	Input_textField_textEdit_2 = ge_go_sdl2.TextField{
		X:           270,
		Y:           40,
		H:           31,
		W:           241,
		Id:          "textEdit_2",
		Size:        10,
		Value:       "",
		BgColor:     DefaultTextFieldColor,
		BorderColor: DefaultTextFieldBorderColor,
		Font:        DefaultFont,
		OnChanged:   textfield_textEdit_2_onchange,
	}
	ge_go_sdl2.HandleAsCallbackArg(textfield_textEdit_2_onchange, OnNewValue_Input_textField_textEdit_2)

	// label_2 text construction
	Input_text_label_2 = ge_go_sdl2.Text{
		Id:      "label_2",
		Content: "Output directory path",
		Font:    DefaultFont,
		Size:    13,
		X:       280,
		Y:       10, W: 231,
		H: 21,

		Alignment: ge_go_sdl2.Left,
		TextColor: sdl.Color{R: 0, G: 0, B: 0, A: 0},
	}
	// pushButton button construction

	// pushButton_label text construction
	Input_text_pushButton_label = ge_go_sdl2.Text{
		Id:      "pushButton_label",
		Content: "",
		Font:    DefaultFont,
		Size:    20,
		X:       0,
		Y:       0, W: 491,
		H: 0,

		Alignment: ge_go_sdl2.Center,
		TextColor: sdl.Color{R: 0, G: 0, B: 0, A: 0},
	}
	Input_button_pushButton_onclick := make(chan string)
	Input_button_pushButton = ge_go_sdl2.Button{
		Id:           "pushButton",
		X:            15,
		Y:            80,
		H:            41,
		W:            491,
		Content:      "Compile",
		BgColor:      DefaultButtonColor,
		BorderColor:  DefaultButtonBorderColor,
		ContentLabel: Input_text_pushButton_label,
		OnClick:      Input_button_pushButton_onclick,
	}
	ge_go_sdl2.HandleAsCallback(Input_button_pushButton_onclick, OnClick_Input_button_pushButton)

	// Input view element construction
	var view_Input_children []interface{}
	view_Input_children = append(view_Input_children, Input_textField_textEdit)
	view_Input_children = append(view_Input_children, Input_text_label)
	view_Input_children = append(view_Input_children, Input_textField_textEdit_2)
	view_Input_children = append(view_Input_children, Input_text_label_2)
	view_Input_children = append(view_Input_children, Input_button_pushButton)
	View_Input = ge_go_sdl2.View{
		X:           0,
		Y:           0,
		H:           136,
		W:           525,
		Id:          "Input",
		BorderColor: DefaultViewBorderColor,
		BgColor:     DefaultViewColor,
		Children:    view_Input_children}

}
