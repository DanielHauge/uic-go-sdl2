package gui

import "github.com/DanielHauge/ge-go-sdl2"

var (
	DefaultFont                        = "D:\\repo\\uic-go-sdl2\\gui\\default.ttf"
	DefaultButtonBorderColor    uint32 = 0x00000000
	DefaultButtonColor          uint32 = 0xD1D1D1D1
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
	construct_View_UICompiler()
	views = append(views, View_UICompiler)
	construct_View_Fail()
	views = append(views, View_Fail)
	construct_View_Input()
	views = append(views, View_Input)
	construct_View_Success()
	views = append(views, View_Success)
	ge_go_sdl2.Run(views)
}
