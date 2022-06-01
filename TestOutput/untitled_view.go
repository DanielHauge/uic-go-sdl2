package gui

import ge_go_sdl2 "github.com/DanielHauge/ge-go-sdl2"

// UI element declaration
var (
	Hello_text_pushButton_label ge_go_sdl2.Text
	Hello_button_pushButton     ge_go_sdl2.Button
	Hello_textField_textEdit    ge_go_sdl2.TextField
	Hello_text_label            ge_go_sdl2.Text
	Hello_container_groupBox    ge_go_sdl2.Container
	View_Hello                  ge_go_sdl2.View
)

// Hello view and elements construction
func construct_View_Hello() {
	// pushButton button construction

	// pushButton_label text construction
	Hello_text_pushButton_label = ge_go_sdl2.Text{
		Id:      "pushButton_label",
		Content: "",
		Font:    DefaultFont,
		Size:    DefaultLabelSize,
		X:       0,
		Y:       0,
	}
	Hello_button_pushButton_onclick := make(chan string)
	Hello_button_pushButton = ge_go_sdl2.Button{
		Id:           "pushButton",
		X:            50,
		Y:            50,
		H:            17,
		W:            56,
		Content:      "PushButton",
		BgColor:      DefaultViewColor,
		BorderColor:  DefaultViewBorderColor,
		ContentLabel: Hello_text_pushButton_label,
		OnClick:      Hello_button_pushButton_onclick,
	}
	ge_go_sdl2.HandleAsCallback(Hello_button_pushButton_onclick, OnClick_Hello_button_pushButton)

	// textEdit text field construction

	textfield_textEdit_onchange := make(chan string)
	Hello_textField_textEdit = ge_go_sdl2.TextField{
		X:           50,
		Y:           100,
		H:           64,
		W:           104,
		Id:          "textEdit",
		Size:        DefaultTextSize,
		Value:       "",
		BgColor:     DefaultViewColor,
		BorderColor: DefaultViewBorderColor,
		Font:        DefaultFont,
		OnChanged:   textfield_textEdit_onchange,
	}
	ge_go_sdl2.HandleAsCallbackArg(textfield_textEdit_onchange, OnNewValue_Hello_textField_textEdit)

	// label text construction
	Hello_text_label = ge_go_sdl2.Text{
		Id:      "label",
		Content: "TextLabel",
		Font:    DefaultFont,
		Size:    DefaultLabelSize,
		X:       240,
		Y:       20,
	}

	// groupBox container construction
	Hello_container_groupBox = ge_go_sdl2.Container{
		X:      60,
		Y:      210,
		Id:     "groupBox",
		ViewId: "GroupBox",
	}

	// Hello view element construction
	var view_Hello_children []interface{}
	view_Hello_children = append(view_Hello_children, Hello_button_pushButton)
	view_Hello_children = append(view_Hello_children, Hello_textField_textEdit)
	view_Hello_children = append(view_Hello_children, Hello_text_label)
	view_Hello_children = append(view_Hello_children, Hello_container_groupBox)
	View_Hello = ge_go_sdl2.View{
		X:           0,
		Y:           0,
		H:           546,
		W:           775,
		Id:          "Hello",
		BorderColor: DefaultViewBorderColor,
		BgColor:     DefaultViewColor,
		Children:    view_Hello_children}

}
