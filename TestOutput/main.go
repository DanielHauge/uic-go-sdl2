package gui

import ge_go_sdl2 "github.com/DanielHauge/ge-go-sdl2"

var (
	DefaultFont                        = "D/repo/uic-go-sdl2/TestOutput/default.ttf"
	DefaultButtonBorderColor    uint32 = 0xffff0000
	DefaultButtonColor          uint32 = 0xffffff00
	DefaultButtonTextSize              = 12
	DefaultLabelSize                   = 14
	DefaultTextSize                    = 14
	DefaultTextFieldColor       uint32 = 0xffffffff
	DefaultTextFieldBorderColor uint32 = 0x00000000
	DefaultViewColor            uint32 = 0xeeeeeeee
	DefaultViewBorderColor      uint32 = 0xffffff00
)

func RunGui() {
	var views []ge_go_sdl2.View
	views = append(views, View_)
	ge_go_sdl2.Run(views)
}
