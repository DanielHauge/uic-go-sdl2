package gui

import "github.com/DanielHauge/ge-go-sdl2"

var (
	DefaultFont                        = "D:\\repo\\uic-go-sdl2\\TestOutput\\default.ttf"
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
	construct_View_MainWindow()
	views = append(views, View_MainWindow)
	construct_View_Hello()
	views = append(views, View_Hello)
	ge_go_sdl2.Run(views)
}
