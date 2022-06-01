package gui

import ge_go_sdl2 "github.com/DanielHauge/ge-go-sdl2"

// UI element declaration
var (
	MainWindow_text_pushButton_label   ge_go_sdl2.Text
	MainWindow_button_pushButton       ge_go_sdl2.Button
	MainWindow_text_pushButton_2_label ge_go_sdl2.Text
	MainWindow_button_pushButton_2     ge_go_sdl2.Button
	View_MainWindow                    ge_go_sdl2.View
)

// MainWindow view and elements construction
func construct_View_MainWindow() {
	// pushButton button construction

	// pushButton_label text construction
	MainWindow_text_pushButton_label = ge_go_sdl2.Text{
		Id:      "pushButton_label",
		Content: "",
		Font:    DefaultFont,
		Size:    DefaultLabelSize,
		X:       0,
		Y:       0,
	}
	MainWindow_button_pushButton_onclick := make(chan string)
	MainWindow_button_pushButton = ge_go_sdl2.Button{
		Id:           "pushButton",
		X:            320,
		Y:            240,
		H:            41,
		W:            141,
		Content:      "PushButton",
		BgColor:      DefaultViewColor,
		BorderColor:  DefaultViewBorderColor,
		ContentLabel: MainWindow_text_pushButton_label,
		OnClick:      MainWindow_button_pushButton_onclick,
	}
	ge_go_sdl2.HandleAsCallback(MainWindow_button_pushButton_onclick, OnClick_MainWindow_button_pushButton)

	// pushButton_2 button construction

	// pushButton_2_label text construction
	MainWindow_text_pushButton_2_label = ge_go_sdl2.Text{
		Id:      "pushButton_2_label",
		Content: "",
		Font:    DefaultFont,
		Size:    DefaultLabelSize,
		X:       0,
		Y:       0,
	}
	MainWindow_button_pushButton_2_onclick := make(chan string)
	MainWindow_button_pushButton_2 = ge_go_sdl2.Button{
		Id:           "pushButton_2",
		X:            250,
		Y:            100,
		H:            17,
		W:            56,
		Content:      "PushButton",
		BgColor:      DefaultViewColor,
		BorderColor:  DefaultViewBorderColor,
		ContentLabel: MainWindow_text_pushButton_2_label,
		OnClick:      MainWindow_button_pushButton_2_onclick,
	}
	ge_go_sdl2.HandleAsCallback(MainWindow_button_pushButton_2_onclick, OnClick_MainWindow_button_pushButton_2)

	// MainWindow view element construction
	var view_MainWindow_children []interface{}
	view_MainWindow_children = append(view_MainWindow_children, MainWindow_button_pushButton)
	view_MainWindow_children = append(view_MainWindow_children, MainWindow_button_pushButton_2)
	View_MainWindow = ge_go_sdl2.View{
		X:           0,
		Y:           0,
		H:           600,
		W:           800,
		Id:          "MainWindow",
		BorderColor: DefaultViewBorderColor,
		BgColor:     DefaultViewColor,
		Children:    view_MainWindow_children}

}
