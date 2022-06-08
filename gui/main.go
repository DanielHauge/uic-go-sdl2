package gui

import "github.com/DanielHauge/ge-go-sdl2"

var (
	DefaultFont                        = "D:\\repo\\uic-go-sdl2\\gui\\default.ttf"
	DefaultButtonBorderColor    uint32 = 0x00000000
	DefaultButtonColor          uint32 = 0xffffffff
	DefaultButtonTextSize              = 12
	DefaultLabelSize                   = 14
	DefaultTextSize                    = 14
	DefaultTextFieldColor       uint32 = 0xffffffff
	DefaultTextFieldBorderColor uint32 = 0x00000000
	DefaultViewColor            uint32 = 0xeeeeeeee
	DefaultViewBorderColor      uint32 = 0x08080808
)

func RunGui() {
	var views []ge_go_sdl2.View
	construct_View_InputView()
	construct_View_MainWindow()
	views = append(views, View_MainWindow)
	views = append(views, View_InputView)

	construct_View_StatusBadView()
	views = append(views, View_StatusBadView)
	construct_View_StatusSuccesView()
	views = append(views, View_StatusSuccesView)
	construct_View_StatusBlankView()
	views = append(views, View_StatusBlankView)
	ge_go_sdl2.Run(views)
}
